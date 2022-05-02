// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"todo/database/ent/task"
	"todo/database/ent/user"
	"todo/database/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescCreatedAt is the schema descriptor for created_at field.
	taskDescCreatedAt := taskFields[3].Descriptor()
	// task.DefaultCreatedAt holds the default value on creation for the created_at field.
	task.DefaultCreatedAt = taskDescCreatedAt.Default.(func() time.Time)
	// taskDescUpdatedAt is the schema descriptor for updated_at field.
	taskDescUpdatedAt := taskFields[4].Descriptor()
	// task.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	task.DefaultUpdatedAt = taskDescUpdatedAt.Default.(func() time.Time)
	// task.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	task.UpdateDefaultUpdatedAt = taskDescUpdatedAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescTaskLimit is the schema descriptor for task_limit field.
	userDescTaskLimit := userFields[2].Descriptor()
	// user.DefaultTaskLimit holds the default value on creation for the task_limit field.
	user.DefaultTaskLimit = userDescTaskLimit.Default.(int)
	// user.TaskLimitValidator is a validator for the "task_limit" field. It is called by the builders before save.
	user.TaskLimitValidator = userDescTaskLimit.Validators[0].(func(int) error)
}
