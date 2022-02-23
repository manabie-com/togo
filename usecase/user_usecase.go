package usecase

import (
	"togo/infrastructure/database"
	"togo/usecase/structs"
)

const TABLENAMEUSER = "users"

type UserUsecase struct {
	database.DbInterface
}

func NewUserUsecase(database database.DbInterface) UserUsecase {
	return UserUsecase{database}
}

func (receiver UserUsecase) Get(getStructRepo structs.GetStruct) structs.GetResultStruct {
	var getResultStructRepo structs.GetResultStruct
	getStructRepo.TableName = TABLENAMEUSER
	getStructInfra := getStructRepo.ConvertGetUsecaseToInfra()
	getResultStructInfra := receiver.DbInterface.Get(getStructInfra)
	getResultStructRepo = getResultStructRepo.ConvertGetResultInfraToUsecase(getResultStructInfra)

	return getResultStructRepo
}
