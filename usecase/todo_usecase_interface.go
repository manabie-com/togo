package usecase

import (
	"togo/usecase/structs"
)

type TodoUsecaseInterface interface {
	Get(getStruct structs.GetStruct) structs.GetResultStruct
}
