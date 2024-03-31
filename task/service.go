package task

type Service interface {
	CreateTask(input Task) (Task, error)
	GetAllTasks() ([]Task, error)
	CompleteSingleTask(input CompleteSingleTaskInput) (CompletedTask, error)
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

func (s *service) CompleteSingleTask(input CompleteSingleTaskInput) (CompletedTask, error) {
	inputTask := CompletedTask{}

	inputTask.TaskID = input.TaskId
	inputTask.CompletedAt = input.CompletedAt

	completedTask, err := s.repository.SaveCompleteTask(inputTask)
	if err != nil {
		return completedTask, err
	}
	return completedTask, nil
}
