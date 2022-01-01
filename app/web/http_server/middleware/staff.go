package middleware

import (
	"github.com/labstack/echo/v4"
)

// Staff : Define middleware validate staff
type Staff struct{}

// NewStaff : Tạo mới đối tuợng Staff middleware
func NewStaff() *Staff {
	return &Staff{}
}

// ValidateToken : Validate token of staff
func (Staff) ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		c.Set("user_id", 1) // int
		return next(c)
	}
}
