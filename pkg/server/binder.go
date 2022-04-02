package server

import "github.com/labstack/echo/v4"

// Binder struct
type Binder struct {
	b echo.Binder
}

// Bind tries to bind request into interface, and if it does then validate it
func (b *Binder) Bind(i interface{}, c echo.Context) error {
	if err := b.b.Bind(i, c); err != nil && err != echo.ErrUnsupportedMediaType {
		return err
	}
	return c.Validate(i)
}

// newBinder initializes custom server binder
func newBinder() *Binder {
	return &Binder{b: &echo.DefaultBinder{}}
}
