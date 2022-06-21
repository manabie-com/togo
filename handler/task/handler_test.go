package task

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_Handler(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockRecordService := new(mockRecordService)
		mockRecordService.On("RecordTask", "1", "todo").Return(nil)
		rr := httptest.NewRecorder()
		router := gin.Default()
		request, err := http.NewRequest(http.MethodPost, "/api/v1/task/record", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		assert.Equal(t, 200, rr.Code)
	})
}