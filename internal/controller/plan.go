package controller

import (
	"fmt"
	"lntvan166/togo/internal/config"
	repo "lntvan166/togo/internal/repository"
	"lntvan166/togo/pkg"
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
	plan, err := repo.Repository.GetPlanByUsername(username)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "failed to get plan!")
		return
	}

	pkg.JSON(w, http.StatusOK, plan)
}

func UpgradePlan(w http.ResponseWriter, r *http.Request) {
	username := context.Get(r, "username").(string)
	admin := config.ADMIN

	if username != admin {
		pkg.ERROR(w, http.StatusBadRequest, fmt.Errorf("you are not admin"), "")
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		pkg.ERROR(w, http.StatusBadRequest, err, "invalid plan id!")
		return
	}

	plan, err := repo.Repository.GetPlanByID(id)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "failed to get plan!")
		return
	}

	if plan == string(vip) {
		pkg.ERROR(w, http.StatusBadRequest, fmt.Errorf("this user have already vip plan"), "")
		return
	}

	err = repo.Repository.UpgradePlan(id, string(vip), config.VIP_LIMIT)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "failed to upgrade plan!")
		return
	}

	pkg.JSON(w, http.StatusOK, "message: upgrade plan success")
}
