package controller

import (
	repo "lntvan166/togo/internal/repository"
	"lntvan166/togo/pkg"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := repo.Repository.GetAllUsers()
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "failed to get users!")
		return
	}

	pkg.JSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		pkg.ERROR(w, http.StatusBadRequest, err, "invalid user id!")
		return
	}

	user, err := repo.Repository.GetUserByID(id)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "failed to get user!")
		return
	}

	pkg.JSON(w, http.StatusOK, user)
}

func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		pkg.ERROR(w, http.StatusBadRequest, err, "invalid user id!")
		return
	}

	err = repo.Repository.DeleteAllTaskOfUser(id)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "failed to delete user!")
		return
	}

	err = repo.Repository.DeleteUserByID(id)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "failed to delete user!")
		return
	}

	pkg.JSON(w, http.StatusOK, "Delete Successfully")
}
