package task

import "gorm.io/gorm"

type Repository interface {
	Save(task Task) (Task, error)
	FindAll() ([]Task, error)
	SaveCompleteTask(completedTask CompletedTask) (CompletedTask, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(task Task) (Task, error) {
	err := r.db.Create(&task).Error
	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *repository) FindAll() ([]Task, error) {
	var tasks []Task
	err := r.db.Preload("CompletedTasks").Find(&tasks).Error
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (r *repository) SaveCompleteTask(completedTask CompletedTask) (CompletedTask, error) {
	err := r.db.Create(&completedTask).Error
	if err != nil {
		return completedTask, err
	}

	return completedTask, nil
}
