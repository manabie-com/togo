package plan

import (
	"fmt"
	"lntvan166/togo/model"
	"lntvan166/togo/utils"
	"net/http"
	"os"
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
	plan, err := model.GetPlanByUsername(username)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf(err.Error()))
		return
	}

	utils.JSON(w, http.StatusOK, plan)
}

func UpgradePlan(w http.ResponseWriter, r *http.Request) {
	username := context.Get(r, "username").(string)
	admin := os.Getenv("ADMIN")

	if username != admin {
		utils.ERROR(w, http.StatusBadRequest, fmt.Errorf("you are not admin"))
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, fmt.Errorf(err.Error()))
		return
	}

	plan, err := model.GetPlanByID(id)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf(err.Error()))
		return
	}

	if plan == string(vip) {
		utils.ERROR(w, http.StatusBadRequest, fmt.Errorf("this user have already vip plan"))
		return
	}

	err = model.UpgradePlan(username, string(vip))
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf(err.Error()))
		return
	}

	utils.JSON(w, http.StatusOK, "message: upgrade plan success")
}
