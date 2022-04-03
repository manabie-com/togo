package subscription

import (
	"net/http"
	"time"

	"github.com/TrinhTrungDung/togo/internal/model"
	"github.com/TrinhTrungDung/togo/pkg/server"
)

func (s *Subscription) Subscribe(authUser *model.AuthUser, data SubscriptionData) error {
	plan := &model.Plan{}
	if err := s.db.Where(&model.Plan{Name: data.PlanName}).Take(plan).Error; err != nil {
		return server.NewHTTPError(http.StatusBadRequest, "INVALID_PLAN", "Plan is currently not supported").SetInternal(err)
	}

	query := &model.Subscription{
		UserID: authUser.ID,
	}
	updates := model.Subscription{
		PlanID: plan.ID,
	}

	if data.PlanName == model.FreemiumPlan {
		updates.EndAt = nil
	} else {
		// Subscription end after 30 days
		endAt := time.Now().Add(time.Duration(720) * time.Hour)
		updates.EndAt = &endAt
	}

	if err := s.db.Where(query).Updates(updates).Error; err != nil {
		return server.NewHTTPInternalError("Cannot update subscription").SetInternal(err)
	}

	return nil
}
