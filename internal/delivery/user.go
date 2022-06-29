package delivery

import (
	"errors"
	"lntvan166/togo/internal/config"
	"lntvan166/togo/internal/domain"
	"lntvan166/togo/pkg"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

type UserDelivery struct {
	UserUsecase domain.UserUsecase
	TaskUsecase domain.TaskUsecase
}

func NewUserDelivery(userUsecase domain.UserUsecase, taskUsecase domain.TaskUsecase) *UserDelivery {
	return &UserDelivery{
		UserUsecase: userUsecase,
		TaskUsecase: taskUsecase,
	}
}

func (u *UserDelivery) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	username := context.Get(r, "username")
	if username != config.ADMIN {
		pkg.ERROR(w, http.StatusUnauthorized, errors.New("authorization failed"), "you are not admin!")
		return
	}

	users, err := u.UserUsecase.GetAllUsers()
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "failed to get users!")
		return
	}

	pkg.JSON(w, http.StatusOK, users)
}

func (u *UserDelivery) GetUser(w http.ResponseWriter, r *http.Request) {
	username := context.Get(r, "username")
	if username != config.ADMIN {
		pkg.ERROR(w, http.StatusUnauthorized, errors.New("authorization failed"), "you are not admin!")
		return
	}

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		pkg.ERROR(w, http.StatusBadRequest, err, "invalid user id!")
		return
	}

	user, err := u.UserUsecase.GetUserByID(id)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "failed to get user!")
		return
	}

	pkg.JSON(w, http.StatusOK, user)
}

func (u *UserDelivery) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	username := context.Get(r, "username")
	if username != config.ADMIN {
		pkg.ERROR(w, http.StatusUnauthorized, errors.New("authorization failed"), "you are not admin!")
		return
	}

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		pkg.ERROR(w, http.StatusBadRequest, err, "invalid user id!")
		return
	}

	err = u.UserUsecase.DeleteUserByID(id)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "failed to delete user!")
		return
	}

	pkg.JSON(w, http.StatusOK, "Delete Successfully")
}
