package usecase

import (
	"togo/infrastructure/database"
	"togo/usecase/structs"
)

const TABLENAMETODOS = "todos"

type TodoUsecase struct {
	database.DbInterface
}

func NewTodoUsecase(database database.DbInterface) TodoUsecase {
	return TodoUsecase{database}
}

func (receiver TodoUsecase) Get(getStructUsecase structs.GetStruct) structs.GetResultStruct {
	var getResultStructUsecase structs.GetResultStruct
	getStructUsecase.TableName = TABLENAMETODOS
	getStructInfra := getStructUsecase.ConvertGetUsecaseToInfra()
	getResultStructInfra := receiver.DbInterface.Get(getStructInfra)
	getResultStructUsecase = getResultStructUsecase.ConvertGetResultInfraToUsecase(getResultStructInfra)

	return getResultStructUsecase
}
