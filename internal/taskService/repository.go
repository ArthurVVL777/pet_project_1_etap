package taskService

import (
	"gorm.io/gorm" // Импортируем библиотеку GORM для работы с базой данных.
)

// TaskRepository определяет методы для работы с задачами в БД.
type TaskRepository interface {
	CreateTask(task Task) (Task, error)              // Метод для создания новой задачи.
	GetAllTasks() ([]Task, error)                    // Метод для получения всех задач.
	UpdateTaskByID(id uint, task Task) (Task, error) // Метод для обновления задачи по ID.
	PatchTaskByID(id uint, task Task) (Task, error)  // Метод для частичного обновления задачи по ID.
	DeleteTaskByID(id uint) error                    // Метод для удаления задачи по ID.
}

// Структура taskRepository реализует интерфейс TaskRepository.
type taskRepository struct {
	db *gorm.DB // Поле db хранит указатель на объект базы данных GORM.
}

// NewTaskRepository создает новый экземпляр репозитория задач.
func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db} // Возвращает новый объект taskRepository с инициализированным полем db.
}

// CreateTask создает новую задачу и сохраняет ее в БД.
func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task) // Используем метод Create GORM для добавления задачи в БД.
	if result.Error != nil {     // Проверяем, произошла ли ошибка при создании.
		return Task{}, result.Error // Возвращаем пустую задачу и ошибку, если она есть.
	}
	return task, nil // Возвращаем созданную задачу и nil как ошибку.
}

// GetAllTasks возвращает все задачи из БД.
func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task               // Объявляем переменную для хранения задач.
	err := r.db.Find(&tasks).Error // Используем метод Find GORM для извлечения всех задач из БД.
	return tasks, err              // Возвращаем массив задач и возможную ошибку.
}

// UpdateTaskByID обновляет задачу по идентификатору.
func (r *taskRepository) UpdateTaskByID(id uint, task Task) (Task, error) {
	var existingTask Task // Объявляем переменную для хранения существующей задачи.

	if err := r.db.First(&existingTask, id).Error; err != nil { // Ищем задачу по ID.
		return Task{}, err // Если задача не найдена или произошла ошибка, возвращаем пустую задачу и ошибку.
	}

	// Обновляем поля существующей задачи новыми значениями из переданного объекта task.
	existingTask.Task = task.Task
	existingTask.IsDone = task.IsDone

	if err := r.db.Save(&existingTask).Error; err != nil { // Сохраняем изменения в БД.
		return Task{}, err // Если произошла ошибка при сохранении, возвращаем пустую задачу и ошибку.
	}
	return existingTask, nil // Возвращаем обновленную задачу и nil как ошибку.
}

// PatchTaskByID частично обновляет задачу по идентификатору.
func (r *taskRepository) PatchTaskByID(id uint, task Task) (Task, error) {
	var existingTask Task // Объявляем переменную для хранения существующей задачи.

	if err := r.db.First(&existingTask, id).Error; err != nil { // Ищем задачу по ID.
		return Task{}, err // Если задача не найдена или произошла ошибка, возвращаем пустую задачу и ошибку.
	}

	// Обновляем только те поля, которые не пустые или истинные в переданном объекте task.
	if task.Task != "" {
		existingTask.Task = task.Task
	}
	if task.IsDone {
		existingTask.IsDone = task.IsDone
	}

	if err := r.db.Save(&existingTask).Error; err != nil { // Сохраняем изменения в БД.
		return Task{}, err // Если произошла ошибка при сохранении, возвращаем пустую задачу и ошибку.
	}
	return existingTask, nil // Возвращаем обновленную задачу и nil как ошибку.
}

// DeleteTaskByID удаляет задачу по идентификатору.
func (r *taskRepository) DeleteTaskByID(id uint) error {
	if err := r.db.Delete(&Task{}, id).Error; err != nil { // Используем метод Delete GORM для удаления задачи по ID.
		return err // Если произошла ошибка при удалении, возвращаем её.
	}
	return nil // Если удаление прошло успешно, возвращаем nil как ошибку.
}
