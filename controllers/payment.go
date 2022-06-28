package controllers

import (
	"database/sql"
	"net/http"

	"github.com/manabie-com/togo/models"
	u "github.com/manabie-com/togo/utils"
)

func Payment(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	decoded := r.Context().Value("user").(*models.Token)
	user := &models.User{
		ID:            decoded.UserId,
		LimitDayTasks: 20,
		IsPayment:     true,
	}

	err := user.UpgradePremium(db)
	if err != nil {
		u.FailureRespond(w, http.StatusBadRequest, "Something went wrong. Please try again")
		return
	}

	u.SuccessRespond(w, http.StatusOK, "Success upgrade Premium account. Please login again to try new upgrade", map[string]interface{}{
		"name":            user.Name,
		"email":           user.Email,
		"is_payment":      user.IsPayment,
		"limit_day_tasks": user.LimitDayTasks,
	})
}
