package delivery

import (
	"fmt"
	"lntvan166/togo/internal/config"
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

func (u *UserDelivery) GetPlan(w http.ResponseWriter, r *http.Request) {
	username := context.Get(r, "username").(string)

	plan, err := u.UserUsecase.GetPlan(username)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "failed to get plan!")
		return
	}

	pkg.JSON(w, http.StatusOK, plan)
}

func (u *UserDelivery) UpgradePlan(w http.ResponseWriter, r *http.Request) {
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

	err = u.UserUsecase.UpgradePlan(id, string(vip), config.VIP_LIMIT)
	if err != nil {
		pkg.ERROR(w, http.StatusInternalServerError, err, "failed to upgrade plan!")
		return
	}

	pkg.JSON(w, http.StatusOK, "message: upgrade plan success")
}
