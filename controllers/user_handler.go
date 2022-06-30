package controllers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"togo-thdung002/entities"
	"togo-thdung002/entities/response"
)

func apiPostUser(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			log.Error("Error on read body request", err)
			return c.JSON(http.StatusBadRequest, resp.Error(err))
		}
		var user entities.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			log.Error("Error on unmarshal body to struct", err)
			return c.JSON(http.StatusInternalServerError, resp.Error(err))
		}
		createdID, err := s.db.CreateUser(&user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, resp.Error(err))
		}
		return c.JSON(http.StatusOK, resp.Success(createdID))
	})
}
