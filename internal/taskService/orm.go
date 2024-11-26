package taskService

// Message представляет простую структуру сообщения.
type Message struct {
	ID   int    `json:"id"`   // Аннотация JSON для идентификатора
	Text string `json:"text"` // Аннотация JSON для текста сообщения
}

// Task представляет задачу с её деталями.
type Task struct {
	ID      uint   `json:"id" gorm:"primaryKey"` // Аннотация JSON и GORM для идентификатора
	Task    string `json:"task"`                 // Аннотация JSON для задачи
	IsDone  bool   `json:"is_done"`              // Аннотация JSON для статуса выполнения
	Message string `json:"message"`              // Аннотация JSON для сообщения
}

// Response представляет структуру ответа для API.
type Response struct {
	Status  string `json:"status"`  // Аннотация JSON для статуса ответа
	Message string `json:"message"` // Аннотация JSON для сообщения ответа
}
