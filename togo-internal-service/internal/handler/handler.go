package handler

import (
	"togo-internal-service/internal/storage"
)

type Handler struct {
	Storage storage.Storage
	Config  *Config
}

type Config struct {
	MaxListTaskPageSize int
}

func (h Handler) Close() error {
	err := h.Storage.Close()
	return err
}

