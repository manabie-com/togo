package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const maxJsonSize = 1024	//1kb

type loginParams struct {
	Id       string	`json:"id"`
	Password string	`json:"password"`
}

func (s *ToDoService) createTokenHandler(resp http.ResponseWriter, req *http.Request) {
	defer func(){
		_ = req.Body.Close()
	}()
	resp.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(io.LimitReader(req.Body, maxJsonSize))

	params := &loginParams{}
	if err := decoder.Decode(params); err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	id := sql.NullString{
		String: params.Id,
		Valid: true,
	}
	pwd := sql.NullString{
		String: params.Password,
		Valid: true,
	}

	if !s.Store.ValidateUser(req.Context(), id, pwd) {
		resp.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect user_id/pwd",
		})
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}

	token, err := s.createToken(id.String)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	err = json.NewEncoder(resp).Encode(map[string]string{
		"data": token,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (s *ToDoService) authHandler(nextHandler http.HandlerFunc) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		req, err := s.validToken(req)
		if err != nil {
			resp.WriteHeader(http.StatusUnauthorized)
			log.Println(err)
			return
		}

		nextHandler(resp, req)
	}
}
