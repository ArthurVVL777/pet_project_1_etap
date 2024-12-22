package taskService

import (
	"gorm.io/gorm"
)

// TaskRepository определяет методы для работы с задачами в БД.
type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskByID(id uint) (Task, error)
	UpdateTaskByID(id uint, task Task) (Task, error)
	PatchTaskByID(id uint, task Task) (Task, error)
	DeleteTaskByID(id uint) error
}

// Структура taskRepository реализует интерфейс TaskRepository.
type taskRepository struct {
	db *gorm.DB
}

// NewTaskRepository создает новый экземпляр репозитория задач.
func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

// Реализация методов интерфейса
func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetTaskByID(id uint) (Task, error) {
	var task Task
	if err := r.db.First(&task, id).Error; err != nil {
		return Task{}, err
	}
	return task, nil
}

func (r *taskRepository) UpdateTaskByID(id uint, task Task) (Task, error) {
	var existingTask Task
	if err := r.db.First(&existingTask, id).Error; err != nil {
		return Task{}, err
	}

	existingTask.Task = task.Task
	existingTask.IsDone = task.IsDone

	if err := r.db.Save(&existingTask).Error; err != nil {
		return Task{}, err
	}

	return existingTask, nil
}

func (r *taskRepository) PatchTaskByID(id uint, task Task) (Task, error) {
	var existingTask Task

	if err := r.db.First(&existingTask, id).Error; err != nil {
		return Task{}, err
	}

	if task.Task != "" {
		existingTask.Task = task.Task
	}

	if task.IsDone != existingTask.IsDone {
		existingTask.IsDone = task.IsDone
	}

	if err := r.db.Save(&existingTask).Error; err != nil {
		return Task{}, err
	}

	return existingTask, nil
}

func (r *taskRepository) DeleteTaskByID(id uint) error {
	if err := r.db.Delete(&Task{}, id).Error; err != nil {
		return err
	}

	return nil
}