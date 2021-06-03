package tasks

import (
	"bytes"
	"net/http"
	"testing"
	"time"

	"github.com/manabie-com/togo/internal/consts"
	"github.com/manabie-com/togo/internal/utils/random"
	"github.com/stretchr/testify/assert"
)

var validCreateRequestJsonString = "{\"content\": \"another content\"}"

func NewMockRequestWithBody(body string) *http.Request {
	bodyByte := []byte(body)
	if body == "" {
		bodyByte = []byte{}
	}
	request, err := http.NewRequest(http.MethodPost, "", bytes.NewReader(bodyByte))
	if err != nil {
		return nil
	}
	return request
}

var CreateRequestBindTestCases = map[string]struct {
	body        string
	expectedErr error
}{
	"Should return error on no body": {
		body:        "",
		expectedErr: consts.ErrInvalidRequest,
	},
	"Should return error on invalid body json": {
		body:        "{",
		expectedErr: consts.ErrInvalidRequest,
	},
	"Should return success non-empty user id and correct body": {
		body:        validCreateRequestJsonString,
		expectedErr: nil,
	},
}

func TestCreateRequestBind(t *testing.T) {
	t.Parallel()

	for tName, tCase := range CreateRequestBindTestCases {
		t.Run(tName, func(t *testing.T) {
			createRequest := &CreateRequest{}
			request := NewMockRequestWithBody(tCase.body)
			got := createRequest.Bind(request)
			assert.Equal(t, tCase.expectedErr, got)
		})
	}
}

func TestCreateRequestToModel(t *testing.T) {
	t.Parallel()

	t.Run("Should return correct model", func(t *testing.T) {
		request := CreateRequest{
			Content:     random.RandString(10),
			CreatedDate: time.Now().Format(consts.DefaultDateFormat),
		}
		userID := random.RandString(10)
		got := request.ToModel(userID)
		assert.Equal(t, request.Content, got.Content)
		assert.Equal(t, request.CreatedDate, got.CreateDate)
		assert.Equal(t, userID, got.UserID)
	})
}
