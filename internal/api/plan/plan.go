package plan

import (
	"github.com/TrinhTrungDung/togo/internal/model"
	"github.com/TrinhTrungDung/togo/pkg/server"
)

var ErrRetrievePlans = server.NewHTTPInternalError("Cannot retrieve plans")

// List returns list of all current provided plans
func (p *Plan) List() ([]*model.Plan, error) {
	var plans []*model.Plan
	if err := p.db.Find(&plans).Error; err != nil {
		return nil, ErrRetrievePlans.SetInternal(err)
	}

	return plans, nil
}
