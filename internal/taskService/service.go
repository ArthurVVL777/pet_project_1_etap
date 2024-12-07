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
func (s *TaskService) UpdateTaskByID(task Task) (Task, error) {
	return s.repo.UpdateTaskByID(task.ID, task)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.DeleteTaskByID(id)
}
