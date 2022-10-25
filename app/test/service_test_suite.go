package test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
	Ctrl *gomock.Controller
}

func (s *ServiceTestSuite) SetupSuite() {
	s.Ctrl = gomock.NewController(s.T())
}

func (s *ServiceTestSuite) TearDownSuite() {
	s.Ctrl.Finish()
}
