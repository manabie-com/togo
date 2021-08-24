package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/internal/dtos"
	"github.com/manabie-com/togo/internal/helpers"
	"github.com/manabie-com/togo/internal/services"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ConfigurationHandler struct {
	configurationService services.ConfigurationService
}

func NewConfigurationHandler(injectedConfigurationService services.ConfigurationService) *ConfigurationHandler {
	return &ConfigurationHandler{
		configurationService: injectedConfigurationService,
	}
}

// GetConfigurationByDate godoc
// @Summary Get Configuration By Date
// @Description Get Configuration By Date
// @Param current_date query string true "Current Date"
// @Tags Configuration
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @Accept  json
// @Produce  json
// @Success 200 {object} dtos.GetListTaskResponse
// @Router /configurations/date [get]
func (h *ConfigurationHandler) GetConfigurationByDate(ctx *gin.Context) {
	currentDate := ctx.Query("current_date")
	userID, ok := helpers.GetUserIdFromContext(ctx)
	if !ok {
		logrus.Errorf("Get Configuration, Get User from context failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	response, err := h.configurationService.GetTaskConfiguration(ctx, userID, currentDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.NewError(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// CreateTaskConfiguration godoc
// @Summary Create New Configuration
// @Description Create New Configuration
// @Param CreateConfigurationRequest body dtos.CreateConfigurationRequest true "Information of CreateConfigurationRequest"
// @Tags Configuration
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @Accept  json
// @Produce  json
// @Success 200 {object} dtos.CreateConfigurationResponse
// @Router /configurations [post]
func (h *ConfigurationHandler) CreateConfiguration(ctx *gin.Context) {
	var request = &dtos.CreateConfigurationRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.NewError(err))
		return
	}

	userID, ok := helpers.GetUserIdFromContext(ctx)
	if !ok {
		logrus.Errorf("Create Configuration Get User from context failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	request.UserID = userID
	response, err := h.configurationService.CreateTaskConfiguration(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.NewError(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
