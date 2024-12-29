package userService

import (
	"gorm.io/gorm"
)

// UserRepository определяет методы для работы с пользователями в БД.
type UserRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	GetUserByID(id uint) (User, error)
	UpdateUserByID(id uint, user User) (User, error)
	DeleteUserByID(id uint) error
}

// userRepository реализует интерфейс UserRepository.
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository создает новый экземпляр репозитория пользователей.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Реализация методов интерфейса
func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	return user, result.Error
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
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

func (r *userRepository) DeleteUserByID(id uint) error {
	return r.db.Delete(&User{}, id).Error
}
