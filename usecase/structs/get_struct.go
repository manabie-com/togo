package structs

import (
	"togo/infrastructure/database/structs"
)

type GetStruct struct {
	TableName  string
	Conditions interface{}
}

func (receiver GetStruct) ConvertGetUsecaseToInfra() structs.GetStruct {
	var getStructInfra structs.GetStruct
	getStructInfra.TableName = receiver.TableName
	getStructInfra.Conditions = receiver.Conditions

	return getStructInfra
}
