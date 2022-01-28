package httpx

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/manabie-com/togo/pkg/errorx"
)

func ParseRequest(r *http.Request, p interface{}) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&p)
	if err != io.EOF {
		return err
	}
	return nil
}

func WriteError(w http.ResponseWriter, err error) {
	errIn := errorx.ToErrorInterface(err)
	jsonErr := errorx.ToErrorJSON(errIn)
	errBody, err := json.Marshal(&jsonErr)
	if err != nil {
		errBody = []byte("{\"type\": \"internal\", \"msg\": \"There was an error but it could not be serialized into JSON\"}") // fallback
	}
	w.Header().Set("Content-Type", "application/json") // Error responses are always JSON
	w.Header().Set("Content-Length", strconv.Itoa(len(errBody)))
	w.WriteHeader(errIn.GetStatusCode()) // set HTTP status code and send response

	w.Write(errBody)
}

func WriteReponse(ctx context.Context, w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`Can not marshal response`))
		return
	}
	w.WriteHeader(status)
	w.Write(response)
}
