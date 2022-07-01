package controllers

import (
	"TOGO/middleware"
	"TOGO/models"
	"TOGO/responses"
	"TOGO/untils"

	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Signup() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		var check models.User
		defer cancel()

		// validate body
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			untils.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		if validationErr := validate.Struct(&user); validationErr != nil {
			untils.Error(rw, validationErr.Error(), http.StatusBadRequest)
			return
		}

		if err := userCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&check); err == nil {
			untils.Error(rw, "User existing", http.StatusBadRequest)
			return
		}

		hashPwd, _ := models.HashPassword(user.Password)

		newUser := models.User{
			Id:       primitive.NewObjectID(),
			Username: user.Username,
			Password: hashPwd,
			Name:     user.Name,
			Limit:    10,
			Status:   true,
			Role:     "user",
		}
		// add obj
		_, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			untils.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		token, err := middleware.CreateToken(newUser.Id, newUser.Role)
		if err != nil {
			untils.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		res := map[string]interface{}{"username": newUser.Username, "name": newUser.Name, "id": newUser.Id}
		responses.WriteResponseUser(rw, token, http.StatusOK, res)
	}
}

func Login() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user_client models.User
		var user models.User
		defer cancel()

		if err := json.NewDecoder(r.Body).Decode(&user_client); err != nil {
			untils.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"username": user_client.Username}).Decode(&user)
		if err != nil {
			untils.Error(rw, "User name or Password Vaild", http.StatusBadRequest)
			return
		}
		// --> user is user sended form client
		if !models.CheckPasswordHash(user_client.Password, user.Password) {
			untils.Error(rw, "Password Vaild", http.StatusBadRequest)
			return
		}

		token, err := middleware.CreateToken(user.Id, user.Role)
		if err != nil {
			untils.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		res := map[string]interface{}{"username": user.Username, "password": user.Password, "id": user.Id}
		responses.WriteResponseUser(rw, token, http.StatusOK, res)

	}
}
