package taskService

import "pet_project_1_etap/internal/userService"

// Message представляет простую структуру сообщения.
type Message struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

// Task представляет задачу с её деталями.
type Task struct {
	ID      uint             `json:"id" gorm:"primaryKey"`
	Task    string           `json:"task"`
	IsDone  bool             `json:"is_done"`
	Message string           `json:"message,omitempty"`
	Text    string           `json:"text,omitempty"`
	UserID  uint             `json:"user_id"` // Новое поле для связи с пользователем
	User    userService.User `json:"user" gorm:"foreignKey:UserID"`
}

// Response представляет структуру ответа для API.
type Response struct {
	Status  string `json:"status"`  // Статус ответа
	Message string `json:"message"` // Сообщение ответа
}

// PostTaskRequestBody определяет тело запроса для создания задачи.
type PostTaskRequestBody struct {
	Task   *string `json:"task"`
	IsDone *bool   `json:"is_done"`
	UserID *uint   `json:"user_id,omitempty"`
}

// PostTasksRequestObject определяет структуру запроса для создания задач.
type PostTasksRequestObject struct {
	Body *PostTaskRequestBody `json:"body"`
}

// PatchTasksJSONRequestBody определяет тело запроса для PATCH задач.
type PatchTasksJSONRequestBody struct {
	Task   *string `json:"task"`
	IsDone *bool   `json:"is_done"`
}

type PatchTasksIdRequestObject struct {
	Id   uint                  `json:"id"`
	Body *PatchTaskRequestBody `json:"body"`
}

type PatchTaskRequestBody struct {
	Task   *string `json:"task"`    // Указатель на текст задачи (опционально)
	IsDone *bool   `json:"is_done"` // Указатель на статус завершенности (опционально)
}
