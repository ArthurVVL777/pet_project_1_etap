package taskService

// TaskService представляет сервис для работы с задачами.
type TaskService struct {
	repo TaskRepository // Репозиторий для работы с задачами в БД.
}

// NewService создает новый экземпляр TaskService с заданным репозиторием.
func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

// CreateTask создает новую задачу через репозиторий.
func (s *TaskService) CreateTask(task Task) (Task, error) {
	return s.repo.CreateTask(task)
}

// GetAllTasks возвращает все задачи через репозиторий.
func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

// GetTaskByID возвращает задачу по идентификатору через репозиторий.
func (s *TaskService) GetTaskByID(id uint) (Task, error) {
	return s.repo.GetTaskByID(id)
}

// PatchTask частично обновляет задачу по идентификатору.
// ID передается как параметр URL, а данные задачи в теле запроса.
func (s *TaskService) PatchTask(id uint, task Task) (Task, error) {
	return s.repo.PatchTaskByID(id, task)
}

// UpdateTaskByID обновляет задачу по идентификатору.
// ID устанавливается перед обновлением задачи.
func (s *TaskService) UpdateTaskByID(id uint, task Task) (Task, error) {
	task.ID = id // Устанавливаем ID задачи перед обновлением
	return s.repo.UpdateTaskByID(id, task)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.DeleteTaskByID(id)
}

func (s *TaskService) GetTasksForUser(userID uint) ([]Task, error) {
	return s.repo.GetTasksByUserID(userID)
}
