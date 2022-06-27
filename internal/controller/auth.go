package controller

import (
	"encoding/json"
	"errors"
	"lntvan166/togo/internal/usecase"
	"lntvan166/togo/pkg"
	"net/http"

	e "lntvan166/togo/internal/entities"

	"github.com/gorilla/context"
)

func Register(w http.ResponseWriter, r *http.Request) {
	user := e.NewUser()
	var err error

	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		pkg.ERROR(w, http.StatusBadRequest, err, "invalid request body!")
		return
	}

	checkUserExist, err := usecase.CheckUserExist(user.Username)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "failed to check user exist!")
		return
	}
	if checkUserExist {
		pkg.ERROR(w, http.StatusBadRequest, errors.New("user already exist"), "")
		return
	}

	user.PreparePassword()

	err = user.IsValid()
	if err != nil {
		pkg.ERROR(w, http.StatusBadRequest, err, "invalid user data!")
		return
	}

	err = usecase.AddUser(user)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "failed to add user!")
		return
	}

	pkg.JSON(w, http.StatusCreated, "Register Successfully")
}

func Login(w http.ResponseWriter, r *http.Request) {
	newUser := e.NewUser()

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		pkg.ERROR(w, http.StatusBadRequest, err, "invalid request body!")
		return
	}

	user, err := usecase.GetUserByName(newUser.Username)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "user not found!")
		return
	}

	if !user.ComparePassWord(newUser.Password) {
		pkg.ERROR(w, http.StatusBadRequest, errors.New("password incorrect"), "")
		return
	}

	token, err := pkg.GenerateToken(user.Username)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "failed to generate token!")
		return
	}

	pkg.JSON(w, http.StatusOK, map[string]string{"token": token, "message": "login successfully"})
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	user, err := usecase.GetUserByName(context.Get(r, "username").(string))
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "failed to get user!")
		return
	}

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		pkg.ERROR(w, http.StatusBadRequest, err, "invalid request body!")
		return
	}

	err = user.IsValid()
	if err != nil {
		pkg.ERROR(w, http.StatusBadRequest, err, "invalid user data!")
		return
	}

	user.PreparePassword()

	err = usecase.UpdateUser(user)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "failed to update user!")
		return
	}

	pkg.JSON(w, http.StatusOK, "Update Password Successfully")
}
