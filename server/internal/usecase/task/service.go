package task

type TaskService struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) Service {
	return &TaskService{
		repo: r,
	}
}
