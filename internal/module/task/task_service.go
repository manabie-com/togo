package task

// Service interface
type Service interface {
	AddTask(userID uint64, content string) (Task, error)
	RetrieveTasks(userID uint64, createdDate string) ([]Task, error)
}

// NewTaskService func
func NewTaskService(repository Repository) (Service, error) {
	return &service{
		repository: repository,
	}, nil
}

type service struct {
	repository Repository
}

func (s *service) AddTask(userID uint64, content string) (Task, error) {
	return s.repository.AddTask(userID, content)
}

func (s *service) RetrieveTasks(userID uint64, createdDate string) ([]Task, error) {
	return s.repository.RetrieveTasks(userID, createdDate)
}
