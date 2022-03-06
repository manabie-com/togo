package codetype

import (
	"strings"

	"github.com/khangjig/togo/util"
)

const PageSizeDefault = 20

type Paginator struct {
	Page  int `json:"page,omitempty" query:"page"`
	Limit int `json:"limit,omitempty" query:"limit"`
}

type GetListRequest struct {
	Paginator Paginator `json:"paginator"`
	SortBy    SortType  `json:"sort_by,omitempty" query:"sort_by"`
	OrderBy   string    `json:"order_by,omitempty" query:"order_by"`
	Search    string    `json:"search,omitempty" query:"search"`
}

func (g *GetListRequest) Format() {
	g.Paginator.Format()
	g.SortBy.Format()
	g.OrderBy = strings.ToLower(strings.TrimSpace(g.OrderBy))
	g.Search = util.SQLEscapeString(strings.TrimSpace(g.Search))
}

func (p *Paginator) Format() {
	if p.Limit == 0 {
		p.Limit = PageSizeDefault
	}

	if p.Limit < 0 {
		p.Limit = -1
	}

	if p.Page <= 0 {
		p.Page = 1
	}
}
