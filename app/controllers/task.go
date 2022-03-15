package controllers

import (
	"strconv"
	"time"
	"togo-service/app/models"
	requests "togo-service/app/request"
	responses "togo-service/app/response"
	"togo-service/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/vcraescu/go-paginator/v2"
	"github.com/vcraescu/go-paginator/v2/adapter"
)

func (handler Handler) CreateTask(ctx *fiber.Ctx) error {
	tokenAuth, err := utils.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	var taskParam requests.CreateTaskParam
	ctx.BodyParser(&taskParam)

	if err := taskParam.Validate(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	var setting models.Setting
	err = handler.DB.Where("user_id = ?", tokenAuth.UserID).First(&setting).Error
	if err != nil || setting.ID == 0 {
		if err != nil {
			println(err.Error())
		}
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid setting. Please contact Administrator",
		})
	}

	var taskModel = models.Task{}
	taskModel.UserID = uint64(tokenAuth.UserID)
	taskModel.Name = taskParam.Name
	taskModel.Description = taskParam.Description
	err = handler.DB.Create(&taskModel).Error

	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	var todayTask int64 = 0
	tz, _ := time.LoadLocation("UTC")
	now := time.Now().In(tz)

	handler.DB.Table("tasks").
		Where("user_id = ? and created_at >= ?", tokenAuth.UserID, now.Format("2006-01-02")).
		Count(&todayTask)

	if todayTask > int64(setting.QuotaPerDay) {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   false,
			"message": "Enough for today. Take a rest!",
		})
	} else {
		taskRes := responses.TaskResponse{
			ID:          taskModel.ID,
			Name:        taskModel.Name,
			Description: taskModel.Description,
			UserID:      taskModel.UserID,
		}

		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
			"error":   false,
			"message": "Created",
			"data":    taskRes,
			"count":   todayTask,
		})
	}
}

func (handler Handler) FetchTasks(ctx *fiber.Ctx) error {
	tokenAuth, err := utils.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	var params requests.PagineUserTaskParam
	ctx.BodyParser(&params)

	if params.Limit == 0 {
		params.Limit = 10
	}

	if params.Limit > 50 { // limit 50 items
		params.Limit = 50
	}

	var items []responses.TaskResponse

	query := handler.DB.Table("tasks").
		Model(models.Task{}).
		Order("created_at desc").
		Where("user_id = ? and deleted_at is NULL", tokenAuth.UserID)

	page := paginator.New(adapter.NewGORMAdapter(query), int(params.Limit))
	totalPage, _ := page.PageNums()
	totalNum, _ := page.Nums()

	if params.Page == 0 || params.Page > totalPage {
		params.Page = 1
	}

	page.SetPage(int(params.Page))
	page.Results(&items)

	res := responses.TaskArrayResponse{
		Items: items,
		MetatData: responses.TaskMetaData{
			Total:       uint(totalNum),
			PageNums:    uint(totalPage),
			PageCurrent: uint(params.Page),
			Limit:       uint(params.Limit),
		},
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (handler Handler) UpdateTask(ctx *fiber.Ctx) error {
	var taskParam requests.UpdateTaskParam
	task_id, _ := strconv.Atoi(ctx.Params("task_id"))
	ctx.BodyParser(&taskParam)
	taskParam.TaskID = task_id

	if err := taskParam.Validate(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	tokenAuth, err := utils.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	var taskModel = models.Task{}
	handler.DB.Find(&taskModel, taskParam.TaskID)

	if taskModel.ID == 0 {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": "Task not found!",
		})
	}

	if !tokenAuth.IsAdmin && taskModel.UserID != uint64(tokenAuth.UserID) {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": "You have not permission to update this task",
		})
	}

	if taskParam.Name != nil {
		taskModel.Name = *taskParam.Name
	}

	if taskParam.Description != nil {
		taskModel.Description = *taskParam.Description
	}

	handler.DB.Save(&taskModel)

	taskRes := responses.TaskResponse{
		ID:          taskModel.ID,
		Name:        taskModel.Name,
		Description: taskModel.Description,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "Created",
		"data":    taskRes,
	})
}

func (handler Handler) DeleteTask(ctx *fiber.Ctx) error {
	task_id, _ := strconv.Atoi(ctx.Params("task_id"))

	tokenAuth, err := utils.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	var taskModel = models.Task{}
	handler.DB.Find(&taskModel, task_id)

	if taskModel.ID == 0 {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": "Task not found!",
		})
	}

	if !tokenAuth.IsAdmin && taskModel.UserID != uint64(tokenAuth.UserID) {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": "You have not permission to delete this task",
		})
	}

	handler.DB.Delete(&taskModel)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "Deleted",
	})
}
