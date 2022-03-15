package structs

import (
	"togo/usecase/structs"
)

type CreateResultStruct struct {
	TableName  string
	Status     string
	Message    string
	Error      error
	Conditions map[string]interface{}
	Data       interface{}
}

func (receiver CreateResultStruct) ConvertCreateResultUsecaseToInterface(getResultStructUsecase structs.CreateResultStruct) CreateResultStruct {
	var createResultStructUsecase CreateResultStruct

	createResultStructUsecase.Status = getResultStructUsecase.Status
	createResultStructUsecase.Message = getResultStructUsecase.Message
	createResultStructUsecase.Error = getResultStructUsecase.Error
	createResultStructUsecase.Data = getResultStructUsecase.Data

	return createResultStructUsecase
}
