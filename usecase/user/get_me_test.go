package user_test

func (suite *TestSuite) TestGetMe_Success() {
	// execute
	_, err := suite.useCase.GetMe(suite.ctx)

	suite.Nil(err)
}
