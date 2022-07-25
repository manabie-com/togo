package render

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/huuthuan-nguyen/manabie/app/transformer"
	"github.com/huuthuan-nguyen/manabie/app/utils"
	validation "github.com/huuthuan-nguyen/manabie/app/validator"
	"net/http"
)

// Error /**
func Error(w http.ResponseWriter, r *http.Request, payload any) {

	var statusCode = http.StatusOK
	var data any
	jsonSerializer := transformer.NewJSONSerializer()

	switch e := payload.(type) {
	case validator.ValidationErrors: // bad request
		statusCode = http.StatusBadRequest
		data = validation.Translate(e)
		jsonSerializer.Messages = []string{http.StatusText(http.StatusNotFound)}
		break
	case utils.Error: // common error
		statusCode = http.StatusBadRequest
		jsonSerializer.Messages = []string{http.StatusText(e.StatusCode)}
		break
	case error: // internal error
		statusCode = http.StatusInternalServerError
		jsonSerializer.Messages = []string{e.Error()}
		break
	default: // default
		statusCode = http.StatusInternalServerError
		jsonSerializer.Messages = []string{http.StatusText(http.StatusInternalServerError)}
	}

	transformerManager := transformer.Manager{
		Serializer: jsonSerializer,
	}
	errorItem := transformer.NewError(data)

	payloadStruct := transformerManager.CreateData(errorItem)
	payloadResponse, err := json.Marshal(payloadStruct)
	if err != nil {
		utils.PanicInternalServerError(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(payloadResponse)
	return
}
