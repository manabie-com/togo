package controllers

import (
	"TOGO/configs"
	"TOGO/models"
	"TOGO/responses"
	"TOGO/untils"
	"fmt"
	"strings"

	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	//"github.com/gorilla/context"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

// var userCollection *dbiface.ConllectionAPI = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func GetUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		userId := mux.Vars(r)["Id"]
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
		if err != nil {
			untils.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		res := map[string]interface{}{"username": user.Username, "id": user.Id}
		responses.WriteResponseUser(rw, "", http.StatusOK, res)
	}
}

func UpdateMe() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		var user_client models.User
		defer cancel()
		// Get id from token
		data := r.Context().Value("Role_Id").(string)
		Role_Id := strings.Split(data, " ")
		// role := Role_Id[0]
		id := Role_Id[1]
		objId, _ := primitive.ObjectIDFromHex(id)

		// r.Body == user_client
		if err := json.NewDecoder(r.Body).Decode(&user_client); err != nil {
			untils.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		//Get User
		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
		if err != nil {
			untils.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		// Check password
		if !models.CheckPasswordHash(user_client.Password, user.Password) {
			untils.Error(rw, "Password Vaild", http.StatusBadRequest)
			return
		}

		update := bson.M{"name": user_client.Name}
		result, err := userCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
		if err != nil {
			untils.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		//get updated task details
		var updatedUser models.User
		if result.MatchedCount == 1 {
			err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)
			if err != nil {
				untils.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		res := map[string]interface{}{"id": updatedUser.Id, "name": updatedUser.Name}
		responses.WriteResponse(rw, http.StatusOK, res)

	}
}

func DeleteUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		// Get role from token
		data := r.Context().Value("Role_Id").(string)
		Role_Id := strings.Split(data, " ")
		role := Role_Id[0]
		defer cancel()
		if role != "admin" {
			untils.Error(rw, "do not permission", http.StatusInternalServerError)
			return
		}

		params := mux.Vars(r)
		userId := params["Id"]
		objId, _ := primitive.ObjectIDFromHex(userId)
		// detele task of user
		results, err := taskCollection.Find(ctx, bson.M{"id_user": objId})
		if err != nil {
			untils.Error(rw, "canot find task of user", http.StatusInternalServerError)
		}
		for results.Next(ctx) {
			var singleTask models.Task
			if err = results.Decode(&singleTask); err != nil {
				untils.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
			task, err := taskCollection.DeleteOne(ctx, bson.M{"id": singleTask.Id})
			if err != nil {
				untils.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
			if task.DeletedCount < 1 {
				untils.Error(rw, "Task with specified ID not found!", http.StatusNotFound)
				return
			}
		}

		// Delete user
		result, err := userCollection.DeleteOne(ctx, bson.M{"id": objId})

		if err != nil {
			untils.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		if result.DeletedCount < 1 {
			untils.Error(rw, "User with specified ID not found!", http.StatusNotFound)
			return
		}

		responses.WriteResponse(rw, http.StatusOK, "Delete completed")
	}
}

func GetAllUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("go to get all")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var users []models.User
		defer cancel()

		// Get role from token
		data := r.Context().Value("Role_Id").(string)
		Role_Id := strings.Split(data, " ")
		role := Role_Id[0]
		//id := Role_Id[1]
		if role != "admin" {
			untils.Error(rw, "do not permission", http.StatusInternalServerError)
			return
		}

		results, err := userCollection.Find(ctx, bson.M{})

		if err != nil {
			untils.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		time.Sleep(time.Second * 2)
		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleUser models.User
			if err = results.Decode(&singleUser); err != nil {
				untils.Error(rw, err.Error(), http.StatusBadRequest)
				return
			}

			users = append(users, singleUser)
		}
		responses.WriteResponse(rw, http.StatusOK, users)
	}
}

func GetMe() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()
		data := r.Context().Value("Role_Id").(string)
		Role_Id := strings.Split(data, " ")
		// role := Role_Id[0]
		id := Role_Id[1]

		objId, _ := primitive.ObjectIDFromHex(id)

		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
		if err != nil {
			untils.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		res := map[string]interface{}{"username": user.Username, "name": user.Name, "id": user.Id, "limit": user.Limit, "role": user.Role}
		responses.WriteResponseUser(rw, "", http.StatusOK, res)

	}
}

func UpdateLimit() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		data := r.Context().Value("Role_Id").(string)
		Role_Id := strings.Split(data, " ")
		//role := Role_Id[0]
		id := Role_Id[1]
		objId, _ := primitive.ObjectIDFromHex(id)

		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
		if err != nil {
			untils.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		update := bson.M{"limit": 100}

		result, err := userCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
		if err != nil {
			untils.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		//get updated task details
		var updatedUser models.User
		if result.MatchedCount == 1 {
			err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)
			if err != nil {
				untils.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		res := map[string]interface{}{"id": updatedUser.Id, "name": updatedUser.Name, "username": updatedUser.Username, "limit": updatedUser.Limit}
		responses.WriteResponse(rw, http.StatusOK, res)
	}
}
