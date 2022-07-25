package handler

import (
	"context"
	"github.com/huuthuan-nguyen/manabie/app/model"
	"github.com/huuthuan-nguyen/manabie/config"
	"github.com/uptrace/bun"
	"log"
)

type Handler struct {
	db     *bun.DB
	ctx    context.Context
	config *config.Config
}

func (handler *Handler) GetDB() *bun.DB {
	return handler.db
}

// New /**
func New(ctx context.Context, config *config.Config) *Handler {

	connection, err := model.NewConnection(config)
	if err != nil {
		log.Fatalf("Error connecting to database...")
	}

	return &Handler{
		db:     connection,
		ctx:    ctx,
		config: config,
	}
}
