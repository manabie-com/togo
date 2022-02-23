package usecase

import (
	"togo/infrastructure/database"
	"togo/usecase/structs"
)

const TABLENAMETODOLIMIT = "todo_limit"

type TodoLimitUsecase struct {
	database.DbInterface
}

func NewTodoLimitUsecase(database database.DbInterface) TodoLimitUsecase {
	return TodoLimitUsecase{database}
}

func (receiver TodoLimitUsecase) Create(createStructUsecase structs.CreateStruct) structs.CreateResultStruct {
	var createResultStructUsecase structs.CreateResultStruct
	createStructUsecase.TableName = TABLENAMETODOLIMIT
	createStructInfra := createStructUsecase.ConvertCreateUsecaseToInfra()
	createResultStructInfra := receiver.DbInterface.Create(createStructInfra)
	createResultStructUsecase = createResultStructUsecase.ConvertCreateResultInfraToUsecase(createResultStructInfra)

	return createResultStructUsecase
}

func (receiver TodoLimitUsecase) Get(getStructUsecase structs.GetStruct) structs.GetResultStruct {
	var getResultStructUsecase structs.GetResultStruct
	getStructUsecase.TableName = TABLENAMETODOLIMIT
	getStructInfra := getStructUsecase.ConvertGetUsecaseToInfra()
	getResultStructInfra := receiver.DbInterface.Get(getStructInfra)
	getResultStructUsecase = getResultStructUsecase.ConvertGetResultInfraToUsecase(getResultStructInfra)

	return getResultStructUsecase
}
