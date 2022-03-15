package usecase

import (
	"togo/usecase/structs"
)

type TodoUsecaseInterface interface {
	Create(createStruct structs.CreateStruct) structs.CreateResultStruct
	Get(getStruct structs.GetStruct) structs.GetResultStruct
}
