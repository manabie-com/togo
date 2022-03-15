package structs

import (
	"togo/usecase/structs"
)

type GetResultStruct struct {
	TableName  string
	Error      error
	Status     string
	Message    string
	Conditions interface{}
	Data       struct {
		RowsAffected int64
		Result       interface{}
	}
}

func (receiver GetResultStruct) ConvertGetResultUsecaseToInterface(getResultStructInfra structs.GetResultStruct) GetResultStruct {
	var getResultStructInterface GetResultStruct

	getResultStructInterface.Error = getResultStructInfra.Error
	getResultStructInterface.Status = getResultStructInfra.Status
	getResultStructInterface.Message = getResultStructInfra.Message
	getResultStructInterface.Conditions = getResultStructInfra.Conditions
	getResultStructInterface.Data = getResultStructInfra.Data

	return getResultStructInterface
}
