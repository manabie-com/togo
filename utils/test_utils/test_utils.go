package test_utils

import (
	"github.com/manabie-com/togo/internal/storages"
)

type RetrievedTasks struct {
	Data []*storages.Task `json:"data"`
}

type CreatedTask struct {
	Data *storages.Task `json:"data"`
}
