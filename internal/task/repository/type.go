package repository

type GetTasksQuery struct {
	UserID int
	Limit  int
	Offset int
}

type GetTaskQuery struct {
	ID int
}
