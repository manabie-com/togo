package handler

import (
	"togo-public-api/internal/service/togo_internal_v1"
	"togo-public-api/internal/service/togo_user_session_v1"
)

type Handler struct {
	TogoInternalService    togo_internal_v1.TogoInternalServiceClient
	TogoUserSessionService togo_user_session_v1.TogoUserSessionServiceClient
}
