package middleware

import (
	"context"
	taskService "github.com/HoangVyDuong/togo/internal/usecase/task"
	userService "github.com/HoangVyDuong/togo/internal/usecase/user"
	"github.com/HoangVyDuong/togo/pkg/define"
	taskDTO "github.com/HoangVyDuong/togo/pkg/dtos/task"
	"github.com/HoangVyDuong/togo/pkg/kit"
	"github.com/HoangVyDuong/togo/pkg/logger"
	"strconv"
)

func LimitCreateTask(taskService taskService.Service, userService userService.Service) kit.Middleware{
	return func(endpoint kit.Endpoint) kit.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			logger.Debug("[RateLimit] Start")
			userID, ok := ctx.Value(define.ContextKeyUserID).(string)
			if !ok {
				logger.Error("[RateLimit] Not found userID from context")
				return nil, define.FailedValidation
			}

			isOver, err := userService.IsOverLimitTask(ctx, userID, define.RateLimitCreateTaskTimes())
			if err != nil {
				logger.Error("[RateLimit] Interact with cache failed")
				return nil, define.Unknown
			}
			if isOver {
				logger.Error("[RateLimit] User Over Limit Task")
				return nil, define.UserOverLimitTask
			}

			resp, err := endpoint(ctx, request)
			if err == nil {
				taskResponse, ok := resp.(taskDTO.CreateTaskResponse)
				taskID, err := strconv.ParseInt(taskResponse.TaskID, 10, 64)
				if err != nil {
					logger.Error("[RateLimit] Response from CreateTask wrong format")
					return nil, define.Unknown
				}
				if !ok {
					logger.Debug("[RateLimit] Not Put On CreateTask")
					return nil, define.Unknown
				}

				taskCreated, err := userService.IncreaseTaskTimesPerDuration(ctx, userID, define.RateLimitCreateTaskDuration())
				if err != nil {
					logger.Errorf("[RateLimit] IncreaseTaskTimePerDay In Cache Failed %s", err.Error())
					if err = taskService.DeleteTask(ctx, taskID); err != nil {
						logger.Errorf("[RateLimit] Rollback Delete Task In DB Failed %s", err.Error())
						return nil, define.Unknown
					}
					return nil, define.Unknown
				}

				if taskCreated > define.RateLimitCreateTaskTimes() {
					logger.Error("[RateLimit] User Over Limit Task")
					if err = taskService.DeleteTask(ctx, taskID); err != nil{
						logger.Errorf("[RateLimit] Rollback Delete Task In DB Failed %s", err.Error())
						return nil, define.Unknown
					}
					return nil, define.UserOverLimitTask
				}
			}

			return response, err
		}
	}
}


