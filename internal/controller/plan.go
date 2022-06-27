package controller

import (
	"fmt"
	"lntvan166/togo/internal/config"
	"lntvan166/togo/internal/repository"
	"lntvan166/togo/internal/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

type Plan string

const (
	free Plan = "free"
	vip  Plan = "vip"
)

func GetPlan(w http.ResponseWriter, r *http.Request) {
	username := context.Get(r, "username").(string)
	plan, err := repository.GetPlanByUsername(username)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "failed to get plan!")
		return
	}

	utils.JSON(w, http.StatusOK, plan)
}

func UpgradePlan(w http.ResponseWriter, r *http.Request) {
	username := context.Get(r, "username").(string)
	admin := config.ADMIN

	if username != admin {
		utils.ERROR(w, http.StatusBadRequest, fmt.Errorf("you are not admin"), "")
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err, "invalid plan id!")
		return
	}

	plan, err := repository.GetPlanByID(id)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "failed to get plan!")
		return
	}

	if plan == string(vip) {
		utils.ERROR(w, http.StatusBadRequest, fmt.Errorf("this user have already vip plan"), "")
		return
	}

	err = repository.UpgradePlan(id, string(vip), config.VIP_LIMIT)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err, "failed to upgrade plan!")
		return
	}

	utils.JSON(w, http.StatusOK, "message: upgrade plan success")
}
