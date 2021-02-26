package services

import (
	"encoding/json"
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
	log.Println(req.Method, req.URL.Path)
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

	usr, err := s.pg.ValidateUser(req.Context(), params.Id, params.Password)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(resp).Encode(newErrResp(err.Error())); err != nil {
			log.Println(err.Error())
		}
		return
	}

	token, err := s.createToken(usr.Id)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(resp).Encode(newErrResp(err.Error())); err != nil {
			log.Println(err.Error())
		}
		return
	}

	if err = json.NewEncoder(resp).Encode(newDataResp(token)); err != nil {
		log.Println(err.Error())
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
