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
		LimitDayTasks: 20,
		IsPayment:     true,
	}

	err := db.QueryRow(`UPDATE users SET is_payment = $1, limit_day_tasks = $2 WHERE id = $3 RETURNING name, email, limit_day_tasks`, user.IsPayment, user.LimitDayTasks, decoded.UserId).Scan(&user.Name, &user.Email, &user.LimitDayTasks)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, "Failure", "Something went wrong. Please try again", nil)
		return
	}

	u.Respond(w, http.StatusOK, "Success", "Success upgrade Premium account. Please login again to try new upgrade", map[string]interface{}{
		"name":            user.Name,
		"email":           user.Email,
		"is_payment":      user.IsPayment,
		"limit_day_tasks": user.LimitDayTasks,
	})
}
