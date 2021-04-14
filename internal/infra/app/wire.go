//+build wireinject

package app

import (
	"context"

	"github.com/google/wire"
)

func InitApplication(ctx context.Context) (*ApplicationContext, func(), error) {
	wire.Build(
		ApplicationSet,
		ApplicationContext{},
	)

	return nil, nil, nil
}
