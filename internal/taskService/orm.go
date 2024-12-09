package taskService

// Message представляет простую структуру сообщения.
type Message struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

// Task представляет задачу с её деталями.
type Task struct {
	ID      uint   //`json:"id" gorm:"primaryKey"`
	Task    string `json:"task"`
	IsDone  bool   `json:"is_done"`
	Message string `json:"message"`
	Text    string
}

// Response представляет структуру ответа для API.
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
type PatchTasksJSONRequestBody struct {
	Task   *string `json:"task"`    // Поле для текста задачи
	IsDone *bool   `json:"is_done"` // Поле для статуса завершенности
}
