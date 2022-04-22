package process

import (
	"ManabieProject/helper"
	dbcontext "ManabieProject/src/dbcontrol"
	"ManabieProject/src/model/dbmodel"
	"ManabieProject/src/model/requestmodel"
	"ManabieProject/src/model/token"
	"errors"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"os"
	"time"
)

// CreateTaskProcess function Create new Task
func CreateTaskProcess(input *requestmodel.CreateTaskRequest) (bool, interface{}) {
	if input != nil {
		if input.Title == "" || input.Details == "" {
			response := helper.Problem{Status: http.StatusBadRequest, Title: http.StatusText(http.StatusBadRequest), Details: "Title or Details is empty !"}
			return false, response
		}

		if input.Token == "" {
			response := helper.Problem{Status: http.StatusNonAuthoritativeInfo, Title: http.StatusText(http.StatusNonAuthoritativeInfo), Details: "Login session expired !"}
			return false, response
		}

		secretKey := os.Getenv("SECRETKEY")
		if secretKey == "" {
			response := helper.Problem{Status: http.StatusInternalServerError, Title: http.StatusText(http.StatusInternalServerError), Details: "Internal Server Error !"}
			return false, response
		}

		// init object JWTMaker
		objectJWTMaker, err := token.NewJWTMaker(secretKey)
		if err != nil {
			response := helper.Problem{Status: http.StatusInternalServerError, Title: http.StatusText(http.StatusInternalServerError), Details: err.Error()}
			return false, response
		}

		payload, err := objectJWTMaker.VerifyToken(input.Token)
		if err != nil || payload == nil {
			response := helper.Problem{Status: http.StatusNonAuthoritativeInfo, Title: http.StatusText(http.StatusNonAuthoritativeInfo), Details: err.Error()}
			return false, response
		}

		idTask, err := createTask(input, payload.Username)
		if err != nil {
			response := helper.Problem{Status: http.StatusLocked, Title: http.StatusText(http.StatusLocked), Details: err.Error()}
			return false, response
		}
		success := helper.Success{Status: http.StatusCreated, Mess: http.StatusText(http.StatusCreated), Data: idTask}
		return true, success
	}
	return false, nil
}

// UpdateTaskProcess update task
func UpdateTaskProcess(input *requestmodel.UpdateTaskRequest) (bool, interface{}) {
	if input != nil {
		if input.Id == "" || input.Details == "" {
			response := helper.Problem{Status: http.StatusBadRequest, Title: http.StatusText(http.StatusBadRequest), Details: "incomplete data !"}
			return false, response
		}

		if input.Status != "" && input.Status != dbmodel.COMPLETE && input.Status != dbmodel.INPROGRESS && input.Status != dbmodel.NOTSATRTED {
			response := helper.Problem{Status: http.StatusBadRequest, Title: http.StatusText(http.StatusBadRequest), Details: "Status can only be complete, inprogress or notstarted !"}
			return false, response
		}

		if input.Token == "" {
			response := helper.Problem{Status: http.StatusNonAuthoritativeInfo, Title: http.StatusText(http.StatusNonAuthoritativeInfo), Details: "Login session expired !"}
			return false, response
		}

		secretKey := os.Getenv("SECRETKEY")
		if secretKey == "" {
			response := helper.Problem{Status: http.StatusInternalServerError, Title: http.StatusText(http.StatusInternalServerError), Details: "Internal Server Error !"}
			return false, response
		}

		// init object JWTMaker
		objectJWTMaker, err := token.NewJWTMaker(secretKey)
		if err != nil {
			response := helper.Problem{Status: http.StatusInternalServerError, Title: http.StatusText(http.StatusInternalServerError), Details: err.Error()}
			return false, response
		}

		payload, err := objectJWTMaker.VerifyToken(input.Token)
		if err != nil || payload == nil {
			response := helper.Problem{Status: http.StatusNonAuthoritativeInfo, Title: http.StatusText(http.StatusNonAuthoritativeInfo), Details: err.Error()}
			return false, response
		}

		idTask, err := updateTask(input, payload.Username)
		if err != nil {
			response := helper.Problem{Status: http.StatusLocked, Title: http.StatusText(http.StatusLocked), Details: err.Error()}
			return false, response
		}
		success := helper.Success{Status: http.StatusCreated, Mess: http.StatusText(http.StatusCreated), Data: idTask}
		return true, success
	}
	return false, nil
}

func createTask(input *requestmodel.CreateTaskRequest, account string) (string, error) {
	if input != nil {
		var accounts = getAccountExists(account)
		if len(accounts) != 1 {
			return "", errors.New("login session expired")
		}
		var account = accounts[0]
		if account.Task == nil {
			account.Task = []dbmodel.Task{}
		}
		if account.CounterTaskPerDay >= account.Group.MaximumTaskPerDay {
			return "", errors.New("maximum task per day")
		}
		var task dbmodel.Task
		task.Id = xid.New().String()
		task.Title = input.Title
		task.Details = input.Details
		task.Status = dbmodel.NOTSATRTED
		task.TaskHistory = []dbmodel.TaskHistory{}
		timed := time.Now()
		task.TaskHistory = append(task.TaskHistory, dbmodel.TaskHistory{Time: &timed, Details: "Create task"})
		account.Task = append(account.Task, task)
		account.CounterTaskPerDay += 1
		// with Account field
		filter := bson.M{
			"$and": []bson.M{
				bson.M{"account": account.Account},
			},
		}
		var db = os.Getenv("DB")
		var collection = os.Getenv("COLLECTION")
		_, err := dbcontext.Context.ReplaceOne(db, collection, filter, account)
		if err != nil {
			return "", err
		}
		return task.Id, nil
	}
	return "", nil
}

func updateTask(input *requestmodel.UpdateTaskRequest, account string) (string, error) {
	if input != nil {
		var accounts = getAccountExists(account)
		if len(accounts) != 1 {
			return "", errors.New("login session expired")
		}
		var account = accounts[0]
		if account.Task == nil {
			account.Task = []dbmodel.Task{}
		}
		for index := range account.Task {
			if input.Id == account.Task[index].Id {
				if input.Status != "" {
					account.Task[index].Status = dbmodel.Status(input.Status)
				} else {
					account.Task[index].Status = dbmodel.INPROGRESS
				}
				time := time.Now()
				taskHistoryNew := dbmodel.TaskHistory{Time: &time, Details: input.Details}
				account.Task[index].TaskHistory = append(account.Task[index].TaskHistory, taskHistoryNew)
			}
		}

		// with Account field
		filter := bson.M{
			"$and": []bson.M{
				bson.M{"account": account.Account},
			},
		}
		var db = os.Getenv("DB")
		var collection = os.Getenv("COLLECTION")
		_, err := dbcontext.Context.ReplaceOne(db, collection, filter, account)
		if err != nil {
			return "", err
		}
		return input.Id, nil
	}
	return "", nil
}
