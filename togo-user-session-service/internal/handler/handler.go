package handler

import (
	"togo-user-session-service/internal/storage"
)

type Handler struct {
	DB storage.Storage
}
