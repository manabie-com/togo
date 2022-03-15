package usecase

import (
	"togo/usecase/structs"
)

type UserUsecaseInterface interface {
	Get(getStruct structs.GetStruct) structs.GetResultStruct
}
