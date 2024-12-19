package taskService

// Message представляет простую структуру сообщения.
type Message struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

// Task представляет задачу с её деталями.
type Task struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Task    string `json:"task"`
	IsDone  bool   `json:"is_done"`
	Message string `json:"message"`
	Text    string `json:"text"`
}

// Response представляет структуру ответа для API.
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// PatchTasksJSONRequestBody определяет тело запроса для PATCH задач.
type PatchTasksJSONRequestBody struct {
	Task   *string `json:"task"`
	IsDone *bool   `json:"is_done"`
	Id     *uint   `json:"id"` // Добавлено поле для ID задачи
}

type PatchTasksRequestObject struct {
	Body *PatchTaskRequestBody `json:"body"`
}

type PatchTaskRequestBody struct {
	Id     *uint   `json:"id"`      // Указатель на ID задачи
	Task   *string `json:"task"`    // Указатель на текст задачи (опционально)
	IsDone *bool   `json:"is_done"` // Указатель на статус завершенности (опционально)
}
