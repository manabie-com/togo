package structs

import (
	"togo/infrastructure/database/structs"
)

type GetResultStruct struct {
	TableName  string
	Error      error
	Status     string
	Message    string
	Conditions map[string]interface{}
	Data       interface{}
}

func (receiver GetResultStruct) ConvertGetResultInfraToUsecase(getResultStructInfra structs.GetResultStruct) GetResultStruct {
	var getResultStructUsecase GetResultStruct

	getResultStructUsecase.Error = getResultStructInfra.Error
	getResultStructUsecase.Status = getResultStructInfra.Status
	getResultStructUsecase.Message = getResultStructInfra.Message
	getResultStructUsecase.Conditions = getResultStructInfra.Conditions
	getResultStructUsecase.Data = getResultStructInfra.Data

	return getResultStructUsecase
}
