package database

import "togo/infrastructure/database/structs"

type DbInterface interface {
	//Create(createStruct CreateStruct) CreateResultStruct
	Get(getStruct structs.GetStruct) structs.GetResultStruct
	//Update(updateStruct UpdateStruct) UpdateResultStruct
	//Delete(deleteStruct DeleteStruct) DeleteResultStruct
}
