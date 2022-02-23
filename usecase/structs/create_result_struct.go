package structs

import (
	"togo/infrastructure/database/structs"
)

type CreateResultStruct struct {
	TableName string
	Status    string
	Message   string
	Error     error
	Data      interface{}
}

func (receiver CreateResultStruct) ConvertCreateResultInfraToUsecase(createResultStructInfra structs.CreateResultStruct) CreateResultStruct {
	var createResultStructRepo CreateResultStruct

	createResultStructRepo.TableName = createResultStructInfra.TableName
	createResultStructRepo.Status = createResultStructInfra.Status
	createResultStructRepo.Message = createResultStructInfra.Message
	createResultStructRepo.Error = createResultStructInfra.Error
	createResultStructRepo.Data = createResultStructInfra.Data

	return createResultStructRepo
}
