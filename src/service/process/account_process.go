package process

import (
	"ManabieProject/helper"
	dbcontext "ManabieProject/src/dbcontrol"
	"ManabieProject/src/model/dbmodel"
	"ManabieProject/src/model/requestmodel"
	"ManabieProject/src/model/token"
	"encoding/base64"
	"go.mongodb.org/mongo-driver/bson"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

// RegisterAccountProcess
func RegisterAccountProcess(input *requestmodel.RegisterRequest) (bool, interface{}) {
	if input != nil {
		if input.Account == "" || input.Password == "" {
			response := helper.Problem{Status: http.StatusBadRequest, Title: http.StatusText(http.StatusBadRequest), Details: "Account or password is empty !"}
			return false, response
		}
		secretKey := os.Getenv("SECRETKEY")
		if secretKey == "" {
			response := helper.Problem{Status: http.StatusInternalServerError, Title: http.StatusText(http.StatusInternalServerError), Details: "Internal Server Error !"}
			return false, response
		}
		// check Account Exists
		if len(getAccountExists(input.Account)) > 0 {
			response := helper.Problem{Status: http.StatusLocked, Title: http.StatusText(http.StatusLocked), Details: "Failed, Account already exists !"}
			return false, response
		}

		err := writeAccountToDb(input)
		if err != nil {
			response := helper.Problem{Status: http.StatusInternalServerError, Title: http.StatusText(http.StatusInternalServerError), Details: err.Error()}
			return false, response
		}

		success := helper.Success{Status: http.StatusCreated, Mess: http.StatusText(http.StatusCreated), Data: nil}

		return true, success
	}
	return false, nil
}

// LoginAccountProcess
func LoginAccountProcess(input *requestmodel.LoginRequest) (bool, interface{}) {
	if input != nil {
		if input.Account == "" || input.Password == "" {
			response := helper.Problem{Status: http.StatusBadRequest, Title: http.StatusText(http.StatusBadRequest), Details: "Account or password is empty !"}
			return false, response
		}
		secretKey := os.Getenv("SECRETKEY")
		if secretKey == "" {
			response := helper.Problem{Status: http.StatusInternalServerError, Title: http.StatusText(http.StatusInternalServerError), Details: "Internal Server Error !"}
			return false, response
		}
		// check Account Exists
		if !validateAccount(input.Account, input.Password) {
			response := helper.Problem{Status: http.StatusBadRequest, Title: http.StatusText(http.StatusBadRequest), Details: "Failed, Account does not exist or wrong password !"}
			return false, response
		}
		// init object JWTMaker
		objectJWTMaker, err := token.NewJWTMaker(secretKey)
		if err != nil {
			response := helper.Problem{Status: http.StatusInternalServerError, Title: http.StatusText(http.StatusInternalServerError), Details: err.Error()}
			return false, response
		}
		duration, _ := strconv.Atoi(os.Getenv("DURATION"))
		durationy := time.Duration(time.Duration(duration) * time.Hour)

		token, err := objectJWTMaker.CreateToken(input.Account, time.Duration(durationy))
		if err != nil {
			response := helper.Problem{Status: http.StatusInternalServerError, Title: http.StatusText(http.StatusInternalServerError), Details: err.Error()}
			return false, response
		}

		success := helper.Success{Status: http.StatusOK, Mess: http.StatusText(http.StatusOK), Data: token}
		return true, success
	}
	return false, nil
}

//checkAccountExists true Exist ,false no exist
func getAccountExists(acc string) []dbmodel.Account {
	var db = os.Getenv("DB")
	var collection = os.Getenv("COLLECTION")

	// with Account field
	filter := bson.D{
		{"account", acc},
	}
	//  option remove id field from all documents
	option := bson.D{{"_id", 0}}

	results, _ := dbcontext.Context.Query(db, collection, filter, option)
	var accounts []dbmodel.Account
	for _, element := range results {
		bsonBytes, err := bson.Marshal(element)
		if err == nil {
			var account dbmodel.Account
			err = bson.Unmarshal(bsonBytes, &account)
			if err == nil {
				accounts = append(accounts, account)
			}
		}
	}
	return accounts
}

// checkAccountExists true is ok  ,false is not ok
func validateAccount(acc string, pass string) bool {
	passBase64 := base64.StdEncoding.EncodeToString([]byte(pass))
	// with Account field
	filter := bson.M{
		"$and": []bson.M{
			bson.M{"account": acc},
			bson.M{"password": passBase64},
		},
	}
	//  option remove id field from all documents
	option := bson.D{{"_id", 0}}

	var db = os.Getenv("DB")
	var collection = os.Getenv("COLLECTION")
	result, err := dbcontext.Context.Query(db, collection, filter, option)
	if err != nil {
		return false
	}
	if len(result) > 0 {
		return true
	}
	return false
}

// writeAccountToDb
func writeAccountToDb(accountRequest *requestmodel.RegisterRequest) error {
	// Convert Data
	var accountInsert dbmodel.Account
	accountInsert.Account = accountRequest.Account
	accountInsert.Password = base64.StdEncoding.EncodeToString([]byte(accountRequest.Password))
	i := rand.Intn(len(dbmodel.GroupDemo))
	accountInsert.Group = dbmodel.GroupDemo[i]
	accountInsert.CounterTaskPerDay = int32(0)
	accountInsert.Task = []dbmodel.Task{}

	var db = os.Getenv("DB")
	var collection = os.Getenv("COLLECTION")
	insertedId, err := dbcontext.Context.InsertOne(db, collection, accountInsert)
	if err != nil || insertedId != 1 {
		return err
	}
	return nil
}
