package routes

import (
	dbcontext "ManabieProject/src/dbcontrol"
	"ManabieProject/src/model/dbmodel"
	function "ManabieProject/src/service/handle"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"time"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	authRoutes := router.Group("api")
	{
		authRoutes.POST("/account/register", function.RegisterAccount)
		authRoutes.POST("/account/login", function.LoginAccount)
		authRoutes.POST("/task/create", function.CreateTask)
		authRoutes.POST("/task/update", function.UpdateTask)
	}
	return router
}

func ResetCounterTask() {
	// Calling NewTicker method
	ticker := time.NewTicker(24 * time.Hour)
	// Creating channel using make
	for {
		select {
		case <-ticker.C:
			resetCounterTask()
		}
	}
}

func resetCounterTask() {
	filter := bson.M{}
	//  option remove id field from all documents
	option := bson.D{{"_id", 0}}
	var db = os.Getenv("DB")
	var collection = os.Getenv("COLLECTION")
	results, err := dbcontext.Context.Query(db, collection, filter, option)
	if err != nil {
		return
	}

	for _, element := range results {
		bsonBytes, err := bson.Marshal(element)
		if err == nil {
			var account dbmodel.Account
			err = bson.Unmarshal(bsonBytes, &account)
			if err == nil {
				filter := bson.M{
					"$and": []bson.M{
						bson.M{"account": account.Account},
					},
				}
				var updateset = bson.D{
					{
						"$set", bson.D{{"counterTaskPerDay", int32(0)}},
					},
				}
				_, _ = dbcontext.Context.UpdateOne(db, collection, filter, updateset)
			}
		}
	}
}
