package id

import "github.com/google/uuid"

type Uid = uuid.UUID

//NewID create a new entity Uid
func NewID() Uid {
	return Uid(uuid.New())
}

func UUIDIsNil(id *uuid.UUID) bool {
	if id == nil {
		return false
	}

	return *id == uuid.Nil
}
