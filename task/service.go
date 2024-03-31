package task

type Service interface {
	CreateTask(input Task) (Task, error)
	GetAllTasks() ([]Task, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateTask(input Task) (Task, error) {

	newTask, err := s.repository.Save(input)
	if err != nil {
		return newTask, err
	}

	return newTask, nil
}

func (s *service) GetAllTasks() ([]Task, error) {
	tasks, err := s.repository.FindAll()
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}
