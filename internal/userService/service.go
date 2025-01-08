package userService

// UserService представляет сервис для работы с пользователями.
type UserService struct {
	repo UserRepository // Репозиторий для работы с пользователями в БД.
}

// NewService создает новый экземпляр UserService с заданным репозиторием.
func NewService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser создает нового пользователя через репозиторий.
func (s *UserService) CreateUser(user User) (User, error) {
	return s.repo.CreateUser(user)
}

// GetAllUsers возвращает всех пользователей через репозиторий.
func (s *UserService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

// GetUserByID возвращает пользователя по идентификатору через репозиторий.
func (s *UserService) GetUserByID(id uint) (User, error) {
	return s.repo.GetUserByID(id)
}

// UpdateUserByID обновляет пользователя по идентификатору через репозиторий.
func (s *UserService) UpdateUserByID(id uint, user User) (User, error) {
	return s.repo.UpdateUserByID(id, user)
}

// PatchUser частично обновляет пользователя по идентификатору через репозиторий.
func (s *UserService) PatchUser(id uint, user User) (User, error) {
	return s.repo.PatchUserByID(id, user)
}

// DeleteUserByID удаляет пользователя по идентификатору через репозиторий.
func (s *UserService) DeleteUserByID(id uint) error {
	return s.repo.DeleteUserByID(id)
}
