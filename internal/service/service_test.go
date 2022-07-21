package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	storemock "github.com/tonghia/togo/internal/store/mocks"
	tasklimitmock "github.com/tonghia/togo/pkg/tasklimit/mocks"
)

type ServiceTestSuite struct {
	suite.Suite
	service          *Service
	mockQuerier      *storemock.MockQuerier
	mockUserLimitSvc *tasklimitmock.MockUserLimitSvc
}

func (suite *ServiceTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(suite.T())
	defer mockCtrl.Finish()

	// store mock
	mockQuerier := storemock.NewMockQuerier(mockCtrl)
	suite.mockQuerier = mockQuerier

	// task limit svc mock
	mockUserLimitSvc := tasklimitmock.NewMockUserLimitSvc(mockCtrl)
	suite.mockUserLimitSvc = mockUserLimitSvc

	service := NewService(mockQuerier, mockUserLimitSvc)
	suite.service = service
}

// Tests

func (suite *ServiceTestSuite) TestNewService() {
	suite.NotNil(suite.service)
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
