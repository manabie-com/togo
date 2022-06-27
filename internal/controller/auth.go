package controller

import (
	"encoding/json"
	"errors"
	"lntvan166/togo/internal/repository"
	"lntvan166/togo/internal/utils"
	"net/http"

	e "lntvan166/togo/internal/entities"

	"github.com/gorilla/context"
)

func Register(w http.ResponseWriter, r *http.Request) {
	user := e.NewUser()
	var err error

	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err, "invalid request body!")
		return
	}

	checkUserExist, err := repository.CheckUserExist(user.Username)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "failed to check user exist!")
		return
	}
	if checkUserExist {
		utils.ERROR(w, http.StatusBadRequest, errors.New("user already exist"), "")
		return
	}

	user.PreparePassword()

	err = user.IsValid()
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err, "invalid user data!")
		return
	}

	err = repository.AddUser(user)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "failed to add user!")
		return
	}

	utils.JSON(w, http.StatusCreated, "Register Successfully")
}

func Login(w http.ResponseWriter, r *http.Request) {
	newUser := e.NewUser()

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err, "invalid request body!")
		return
	}

	user, err := repository.GetUserByName(newUser.Username)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "user not found!")
		return
	}

	if !user.ComparePassWord(newUser.Password) {
		utils.ERROR(w, http.StatusBadRequest, errors.New("password incorrect"), "")
		return
	}

	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "failed to generate token!")
		return
	}

	utils.JSON(w, http.StatusOK, map[string]string{"token": token, "message": "login successfully"})
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	user, err := repository.GetUserByName(context.Get(r, "username").(string))
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "failed to get user!")
		return
	}

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err, "invalid request body!")
		return
	}

	err = user.IsValid()
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err, "invalid user data!")
		return
	}

	user.PreparePassword()

	err = repository.UpdateUser(user)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "failed to update user!")
		return
	}

	utils.JSON(w, http.StatusOK, "Update Password Successfully")
}
