package handlers

import (
	"context"
	"pet_project_1_etap/internal/taskService"
	"pet_project_1_etap/internal/web/tasks"
)

type Handler struct {
	Service *taskService.TaskService
}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	// Получение всех задач из сервиса
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := tasks.GetTasks200JSONResponse{}

	// Заполняем слайс response всеми задачами из БД
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Text,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := taskService.Task{
		Text:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Text,
		IsDone: &createdTask.IsDone,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func (h *Handler) PatchTasks(ctx context.Context, request tasks.PatchTasksRequestObject) (tasks.PatchTasksResponseObject, error) {
	// Извлекаем идентификатор задачи из запроса
	taskID := request.Path.Id
	// Извлекаем тело запроса, которое содержит обновленные данные задачи
	taskRequest := request.Body

	// Создаем объект задачи для обновления
	taskToUpdate := taskService.Task{
		ID:     taskID,
		Text:   taskRequest.Task,
		IsDone: taskRequest.IsDone,
	}

	// Обращаемся к сервису для обновления задачи
	updatedTask, err := h.Service.UpdateTask(taskToUpdate)
	if err != nil {
		return nil, err // Возвращаем ошибку, если обновление не удалось
	}

	// Создаем структуру ответа с обновленными данными задачи
	response := tasks.PatchTasks200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Text,
		IsDone: &updatedTask.IsDone,
	}

	return response, nil // Возвращаем ответ с обновленной задачей
}

func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	// Извлекаем идентификатор задачи из запроса
	taskID := request.Path.Id

	// Вызываем метод сервиса для удаления задачи
	err := h.Service.DeleteTask(taskID)
	if err != nil {
		return nil, err // Возвращаем ошибку, если удаление не удалось
	}

	// Возвращаем успешный ответ без тела
	return tasks.DeleteTasksId204ResponseObject{}, nil
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}
