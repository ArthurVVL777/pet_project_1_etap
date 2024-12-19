package taskService

// Message представляет простую структуру сообщения.
type Message struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

// Task представляет задачу с её деталями.
type Task struct {
	ID      uint   `json:"id" gorm:"primaryKey"` // ID задачи, ключ в базе данных
	Task    string `json:"task"`                 // Описание задачи
	IsDone  bool   `json:"is_done"`              // Статус завершенности задачи
	Message string `json:"message"`              // Дополнительное сообщение (если нужно)
	Text    string `json:"text"`                 // Дополнительный текст (если нужно)
}

// Response представляет структуру ответа для API.
type Response struct {
	Status  string `json:"status"`  // Статус ответа
	Message string `json:"message"` // Сообщение ответа
}

// PostTaskRequestBody определяет тело запроса для создания задачи.
type PostTaskRequestBody struct {
	Task   *string `json:"task"`    // Указатель на текст задачи (опционально)
	IsDone *bool   `json:"is_done"` // Указатель на статус завершенности (опционально)
}

// PostTasksRequestObject определяет структуру запроса для создания задач.
type PostTasksRequestObject struct {
	Body *PostTaskRequestBody `json:"body"`
	Id   *uint                `json:"id"` // Указатель на ID задачи (опционально)
}

// PatchTasksJSONRequestBody определяет тело запроса для PATCH задач.
type PatchTasksJSONRequestBody struct {
	Task   *string `json:"task"`    // Указатель на текст задачи (опционально)
	IsDone *bool   `json:"is_done"` // Указатель на статус завершенности (опционально)
	Id     *uint   `json:"id"`      // Указатель на ID задачи (опционально)
}

// PatchTasksRequestObject определяет структуру запроса для обновления задач.
type PatchTasksRequestObject struct {
	Body *PatchTaskRequestBody `json:"body"`
}

// PatchTaskRequestBody определяет тело запроса для частичного обновления задач.
type PatchTaskRequestBody struct {
	Id     *uint   `json:"id"`      // Указатель на ID задачи
	Task   *string `json:"task"`    // Указатель на текст задачи (опционально)
	IsDone *bool   `json:"is_done"` // Указатель на статус завершенности (опционально)
}
