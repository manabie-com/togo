package integrationtest

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func (s *e2eTestSuite) Test_E2E_CreateTask_ErrInvalidToken() {
	req, err := http.NewRequest(
		"POST",
		s.getClientURL("tasks"),
		strings.NewReader(`{"title":"task 1", "description": "description 1"}`),
	)
	s.NoError(err)

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := http.Client{}
	response, err := client.Do(req)
	defer response.Body.Close()
	s.NoError(err)
	s.Equal(http.StatusBadRequest, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)
	s.Equal(
		`{"status_code":400,"message":"wrong auth header","log":"wrong auth header","error_key":"ErrWrongAuthHeader"}`,
		strings.Trim(string(byteBody), "\n"),
	)
}
