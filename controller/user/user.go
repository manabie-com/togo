package user

import (
	"lntvan166/togo/model"
	"lntvan166/togo/utils"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := model.GetAllUsers()
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "failed to get users!")
		return
	}

	utils.JSON(w, http.StatusOK, users)
}
