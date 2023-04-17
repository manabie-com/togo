package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	router := gin.Default()
	router.GET("/ping", ping())

	req, _ := http.NewRequest("GET", "/ping", bytes.NewBuffer(nil))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	var r map[string]any
	json.Unmarshal(w.Body.Bytes(), &r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, r)
}
