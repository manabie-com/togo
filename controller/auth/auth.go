package auth

import (
	"encoding/json"
	"errors"
	"lntvan166/togo/model"
	"lntvan166/togo/utils"
	"net/http"

	e "lntvan166/togo/entities"
)

func Register(w http.ResponseWriter, r *http.Request) {
	user := e.NewUser()

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if model.CheckUserExist(user.Username) {
		utils.ERROR(w, http.StatusBadRequest, errors.New("user already exist"))
		return
	}

	user.PreparePassword()

	err = user.IsValid()
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = model.AddUser(user)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, 201, "Register Successfully")
}

func Login(w http.ResponseWriter, r *http.Request) {
	newUser := e.NewUser()

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user, err := model.GetUserByName(newUser.Username)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, errors.New("user not found"))
		return
	}

	if !user.ComparePassWord(newUser.Password) {
		utils.ERROR(w, http.StatusBadRequest, errors.New("password incorrect"))
		return
	}

	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, map[string]string{"token": token, "message": "login successfully"})
}
