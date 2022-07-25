package render

import (
	"encoding/json"
	"github.com/huuthuan-nguyen/manabie/app/transformer"
	"github.com/huuthuan-nguyen/manabie/app/utils"
	"net/http"
)

// JSON /**
func JSON(w http.ResponseWriter, r *http.Request, payload interface{}) {
	transformerManager := transformer.NewManager()

	payloadStruct := transformerManager.CreateData(payload)
	payloadResponse, err := json.Marshal(payloadStruct)
	if err != nil {
		utils.PanicInternalServerError(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(payloadResponse)
}
