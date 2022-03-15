package structs

import "togo/infrastructure/database/structs"

type CreateStruct struct {
	TableName string
	Data      interface{}
}

func (receiver CreateStruct) ConvertCreateUsecaseToInfra() structs.CreateStruct {
	var createStructInfra structs.CreateStruct
	createStructInfra.TableName = receiver.TableName
	createStructInfra.Data = receiver.Data

	return createStructInfra
}
