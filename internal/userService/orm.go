package userService

import (
	"pet_project_1_etap/internal/taskService"
	"time"
)

// User представляет пользователя с его данными.
type User struct {
	ID        uint               `json:"id" gorm:"primaryKey"`
	Email     string             `json:"email" gorm:"unique;not null"`
	Password  string             `json:"password" gorm:"not null"`
	CreatedAt time.Time          `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time          `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt *time.Time         `json:"deletedAt" gorm:"index"`
	Tasks     []taskService.Task `json:"tasks" gorm:"foreignKey:UserID"`
}

// Response представляет структуру ответа для API.
type Response struct {
	Status  string `json:"status"`  // Статус ответа
	Message string `json:"message"` // Сообщение ответа
}

// PostUserRequestBody определяет тело запроса для создания пользователя.
type PostUserRequestBody struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

// PostUsersRequestObject определяет структуру запроса для создания пользователей.
type PostUsersRequestObject struct {
	Body *PostUserRequestBody `json:"body"`
}

// PatchUserJSONRequestBody определяет тело запроса для PATCH пользователей.
type PatchUserJSONRequestBody struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

type PatchUserIdRequestObject struct {
	Id   uint                  `json:"id"`
	Body *PatchUserRequestBody `json:"body"`
}

type PatchUserRequestBody struct {
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}
