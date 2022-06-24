package user

import (
	"lntvan166/togo/model"
	"lntvan166/togo/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := model.GetAllUsers()
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "failed to get users!")
		return
	}

	utils.JSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err, "invalid user id!")
		return
	}

	user, err := model.GetUserByID(id)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "failed to get user!")
		return
	}

	utils.JSON(w, http.StatusOK, user)
}

func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err, "invalid user id!")
		return
	}

	err = model.DeleteAllTaskOfUser(id)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "failed to delete user!")
		return
	}

	err = model.DeleteUserByID(id)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "failed to delete user!")
		return
	}

	utils.JSON(w, http.StatusOK, "Delete Successfully")
}
