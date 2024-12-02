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

func (s *TaskService) PatchTask(id uint, task Task) (Task, error) {
	return s.repo.PatchTaskByID(id, task)
}
func (s *TaskService) UpdateTaskByID(id uint, task Task) (Task, error) {
	return s.repo.UpdateTaskByID(id, task)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.DeleteTaskByID(id)
}

// GetAllTasks возвращает все задачи из репозитория.
func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

// PatchTask обновляет существующую задачу по ID.
func (s *TaskService) PatchTask(id uint, task Task) (Task, error) {
	return s.repo.PatchTaskByID(id, task)
}

// DeleteTask удаляет задачу по ID.
func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.DeleteTaskByID(id)
}
