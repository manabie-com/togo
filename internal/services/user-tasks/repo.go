package user_tasks

import "github.com/looplab/eventhorizon"

type UserTaskRepo interface {
	eventhorizon.ReadWriteRepo
}
