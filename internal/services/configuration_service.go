package services

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/manabie-com/togo/internal/dtos"
	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/repositories"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type configurationService struct {
	configurationRepository repositories.ConfigurationRepository
}

type ConfigurationService interface {
	CreateTaskConfiguration(ctx context.Context, request *dtos.CreateConfigurationRequest) (*dtos.CreateConfigurationResponse, error)
	GetTaskConfiguration(ctx context.Context, userID string, date string) (int64, error)
}

func NewConfigurationService(injectedConfigurationRepository repositories.ConfigurationRepository) ConfigurationService {
	return &configurationService{configurationRepository: injectedConfigurationRepository}
}

func (s *configurationService) CreateTaskConfiguration(ctx context.Context, request *dtos.CreateConfigurationRequest) (*dtos.CreateConfigurationResponse, error) {
	var configuration = &models.Configuration{
		ID:       uuid.New().String(),
		UserID:   request.UserID,
		Capacity: request.Capacity,
		Date:     request.Date,
	}

	createdConfiguration, err := s.configurationRepository.CreateConfiguration(ctx, configuration)
	if err != nil {
		logrus.Errorf("Create Configuration error: %s", err.Error())
		return nil, err
	}

	var response = &dtos.ConfigurationDto{}
	if err := copier.Copy(response, createdConfiguration); err != nil {
		logrus.Errorf("Create Configuration Mapping Dto error: %s", err.Error())
		return nil, err
	}

	return &dtos.CreateConfigurationResponse{Data: response}, nil
}

func (s *configurationService) GetTaskConfiguration(ctx context.Context, userID string, date string) (int64, error) {
	capacity, err := s.configurationRepository.GetCapacity(ctx, userID, date)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("Get Configuration error: %s", err.Error())
		return 0, err
	}

	return capacity, nil
}
