package transport

import (
	"github.com/manabie-com/togo/internal/usecase"
)

//TogoHandler represent the httphandler for togo
type TogoHandler struct {
	TogoUsecase usecase.TogoUsecase
}
