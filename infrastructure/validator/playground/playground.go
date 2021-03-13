package playground

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-playground/validator"

	_validator "github.com/valonekowd/togo/infrastructure/validator"
)

type playgroundValidator struct {
	logger    log.Logger
	validator *validator.Validate
}

func NewValidator(tagName string, logger log.Logger) _validator.Validator {
	v := &playgroundValidator{
		logger:    logger,
		validator: validator.New(),
	}

	v.setup(tagName)

	return v
}

func (p *playgroundValidator) setup(tagName string) {
	p.validator.SetTagName(tagName)
}

func (p *playgroundValidator) Struct(ctx context.Context, s interface{}) error {
	return p.validator.StructCtx(ctx, s)
}
