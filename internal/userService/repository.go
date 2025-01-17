package userService

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"pet_project_1_etap/internal/taskService"
)

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	GetUserByID(id uint) (User, error)
	UpdateUserByID(id uint, user User) (User, error)
	PatchUserByID(id uint, user User) (User, error) // Метод для частичного обновления
	DeleteUserByID(id uint) error
	GetTasksForUser(id uint) ([]taskService.Task, error)
}

// userRepository реализует интерфейс UserRepository.
type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) GetTasksForUser(userID uint) ([]taskService.Task, error) {
	var tasks []taskService.Task
	if err := r.db.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// NewUserRepository создает новый экземпляр репозитория пользователей.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	return user, result.Error
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var usersList []User
	err := r.db.Find(&usersList).Error
	return usersList, err
}

func (r *userRepository) GetUserByID(id uint) (User, error) {
	var user User
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *userRepository) UpdateUserByID(id uint, user User) (User, error) {
	var existingUser User
	if err := r.db.First(&existingUser, id).Error; err != nil {
		return User{}, err
	}

	existingUser.Email = user.Email
	existingUser.Password = user.Password

	if err := r.db.Save(&existingUser).Error; err != nil {
		return User{}, err
	}

	return existingUser, nil
}

// PatchUserByID частично обновляет пользователя по идентификатору.
func (r *userRepository) PatchUserByID(id uint, user User) (User, error) {
	var existingUser User
	if err := r.db.First(&existingUser, id).Error; err != nil {
		return User{}, err
	}

	// Обновляем только те поля, которые были переданы
	if user.Email != "" {
		existingUser.Email = user.Email
	}

	if user.Password != "" {
		existingUser.Password = user.Password
	}

	if err := r.db.Save(&existingUser).Error; err != nil {
		return User{}, err
	}

	return existingUser, nil
}

func (r *userRepository) DeleteUserByID(id uint) error {
	return r.db.Delete(&User{}, id).Error
}

// ServerInterface определяет методы для работы с пользователями.
type ServerInterface interface {
	GetUsers(ctx echo.Context) error
	PostUsers(ctx echo.Context) error
	PatchUsersId(ctx echo.Context, id uint) error
	DeleteUsersId(ctx echo.Context, id uint) error
	GetUsersUserIdTasks(ctx echo.Context, userId uint) error
}
