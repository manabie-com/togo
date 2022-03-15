package structs

import (
	"togo/usecase/structs"
)

type CreateStruct struct {
	TableName string
	Data      interface{}
}

func (receiver CreateStruct) ConvertCreateInterfaceToUsecase() structs.CreateStruct {
	var createStructInfra structs.CreateStruct
	createStructInfra.TableName = receiver.TableName
	createStructInfra.Data = receiver.Data

	return createStructInfra
}
