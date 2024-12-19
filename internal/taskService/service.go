package taskService

type TaskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task Task) (Task, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) GetTaskByID(id uint) (Task, error) {
	return s.repo.GetTaskByID(id)
}

// Изменяем метод PatchTask для поддержки обновления по ID
func (s *TaskService) PatchTask(id uint, task Task) (Task, error) {
	return s.repo.PatchTaskByID(id, task)
}

// Обновляем метод UpdateTaskByID для поддержки передачи ID
func (s *TaskService) UpdateTaskByID(id uint, task Task) (Task, error) {
	task.ID = id // Устанавливаем ID задачи перед обновлением
	return s.repo.UpdateTaskByID(id, task)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.DeleteTaskByID(id)
}
