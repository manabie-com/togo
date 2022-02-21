package structs

import (
	"togo/usecase/structs"
)

type GetStruct struct {
	TableName  string
	Conditions map[string]interface{}
}

func (receiver GetStruct) ConvertGetInterfaceToUsecase() structs.GetStruct {
	var getStructInfra structs.GetStruct
	getStructInfra.TableName = receiver.TableName
	getStructInfra.Conditions = receiver.Conditions

	return getStructInfra
}
