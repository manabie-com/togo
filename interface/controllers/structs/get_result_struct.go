package structs

import (
	"togo/usecase/structs"
)

type GetResultStruct struct {
	TableName  string
	Error      error
	Status     string
	Message    string
	Conditions map[string]interface{}
	Data       interface{}
}

func (receiver GetResultStruct) ConvertGetResultUsecaseToInterface(getResultStructInfra structs.GetResultStruct) GetResultStruct {
	var getResultStructRepo GetResultStruct

	getResultStructRepo.Error = getResultStructInfra.Error
	getResultStructRepo.Status = getResultStructInfra.Status
	getResultStructRepo.Message = getResultStructInfra.Message
	getResultStructRepo.Conditions = getResultStructInfra.Conditions
	getResultStructRepo.Data = getResultStructInfra.Data

	return getResultStructRepo
}
