package task

type TaskMapper struct {
}

func (tm *TaskMapper) ToDatabase(entity *TaskMapper) (interface{}, error) {
	return nil, nil
}

func (tm *TaskMapper) ToEntity(database interface{}) (*TaskMapper, error) {
	return nil, nil
}
