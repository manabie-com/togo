package integrationtest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (s *e2eTestSuite) Test_E2E_RegisterSucceed() {
	req, err := http.NewRequest(
		"POST",
		s.getClientURL("register"),
		strings.NewReader(`{"email":"register@gmail.com", "password": "register@123"}`),
	)
	s.NoError(err)

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := http.Client{}
	response, err := client.Do(req)
	defer response.Body.Close()
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(byteBody))
	s.NoError(err)
	s.NotNil(strings.Trim(string(byteBody), ""))
}

func (s *e2eTestSuite) Test_E2E_LoginSucceed() {

	req, err := http.NewRequest(
		http.MethodPost,
		s.getClientURL("login"),
		strings.NewReader(`{"email":"login@gmail.com", "password": "login@123"}`),
	)
	s.NoError(err)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := http.Client{}
	res, err := client.Do(req)
	defer res.Body.Close()
	s.NoError(err)
	s.Equal(http.StatusOK, res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	s.NoError(err)
	s.NotNil(strings.Trim(string(body), ""))
}

//func (s *e2eTestSuite) Test_E2E_LoginFail_WrongEmail() {
//	req, err := http.NewRequest(
//		"POST",
//		s.getClientURL("login"),
//		strings.NewReader(`{"email":"user1@gmail.com", "password": "user@123"}`),
//	)
//	s.NoError(err)
//
//	req.Header.Set("Content-Type", "application/json; charset=utf-8")
//
//	client := http.Client{}
//	response, err := client.Do(req)
//	defer response.Body.Close()
//	s.NoError(err)
//	s.Equal(http.StatusBadRequest, response.StatusCode)
//
//	byteBody, err := ioutil.ReadAll(response.Body)
//	s.NoError(err)
//	s.Equal(
//		`{"status_code":400,"message":"email or password invalid","log":"email or password invalid","error_key":"ErrEmailOrPasswordInvalid"}`,
//		strings.Trim(string(byteBody), "\n"),
//	)
//}
//
//func (s *e2eTestSuite) Test_E2E_LoginFail_WrongPassword() {
//	req, err := http.NewRequest(
//		"POST",
//		s.getClientURL("login"),
//		strings.NewReader(`{"email":"user@gmail.com", "password": "user@1234"}`),
//	)
//	s.NoError(err)
//
//	req.Header.Set("Content-Type", "application/json; charset=utf-8")
//
//	client := http.Client{}
//	response, err := client.Do(req)
//	defer response.Body.Close()
//	s.NoError(err)
//	s.Equal(http.StatusBadRequest, response.StatusCode)
//
//	byteBody, err := ioutil.ReadAll(response.Body)
//	s.NoError(err)
//	s.Equal(
//		`{"status_code":400,"message":"email or password invalid","log":"email or password invalid","error_key":"ErrEmailOrPasswordInvalid"}`,
//		strings.Trim(string(byteBody), "\n"),
//	)
//}
//
//func (s *e2eTestSuite) Test_E2E_LoginFail_MissingEmail() {
//	req, err := http.NewRequest(
//		"POST",
//		s.getClientURL("login"),
//		strings.NewReader(`{"password": "user@123"}`),
//	)
//	s.NoError(err)
//
//	req.Header.Set("Content-Type", "application/json; charset=utf-8")
//
//	client := http.Client{}
//	response, err := client.Do(req)
//	defer response.Body.Close()
//	s.NoError(err)
//	s.Equal(http.StatusBadRequest, response.StatusCode)
//
//	byteBody, err := ioutil.ReadAll(response.Body)
//	s.NoError(err)
//	s.Equal(
//		`{"status_code":400,"message":"invalid request","log":"Key: 'UserLogin.Email' Error:Field validation for 'Email' failed on the 'required' tag","error_key":"ErrInvalidRequest"}`,
//		strings.Trim(string(byteBody), "\n"),
//	)
//}
