package handlers

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"pet_project_1_etap/internal/taskService"
	"pet_project_1_etap/internal/web/tasks"
)

type Handler struct {
	Service *taskService.TaskService
}

// GetTasks обрабатывает запрос на получение всех задач.
func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		log.Printf("Error fetching tasks: %v", err)
		return nil, err
	}

	response := make(tasks.GetTasks200JSONResponse, 0, len(allTasks))
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Text,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}

// PostTasks обрабатывает запрос на создание новой задачи.
func (h *Handler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body
	if taskRequest == nil || taskRequest.Task == nil || taskRequest.IsDone == nil {
		log.Printf("Request body or fields cannot be nil")
		return nil, fmt.Errorf("request body or fields cannot be nil")
	}

	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}

	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		log.Printf("Error creating task: %v", err)
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}

	return response, nil
}

// PatchTasksId обрабатывает запрос на обновление существующей задачи.
// PatchTasksId обрабатывает запрос на обновление существующей задачи.
func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	if request.Body.Id == nil {
		log.Printf("Task ID cannot be nil")
		return nil, fmt.Errorf("task ID cannot be nil")
	}

	taskID := *request.Body.Id // Разыменовываем указатель на ID задачи
	taskRequest := request.Body

	if taskRequest == nil {
		log.Printf("Request body cannot be nil")
		return nil, fmt.Errorf("request body cannot be nil")
	}

	existingTask, err := h.Service.GetTaskByID(taskID)
	if err != nil {
		log.Printf("Error fetching task with ID %d: %v", taskID, err)
		return nil, fmt.Errorf("task not found: %v", err)
	}

	if taskRequest.Task != nil {
		existingTask.Task = *taskRequest.Task // Обновляем текст задачи только если он не nil
	}

	if taskRequest.IsDone != nil {
		existingTask.IsDone = *taskRequest.IsDone // Обновляем статус завершенности только если он не nil
	}

	updatedTask, err := h.Service.UpdateTaskByID(taskID, existingTask)
	if err != nil {
		log.Printf("Error updating task with ID %d: %v", taskID, err)
		return nil, fmt.Errorf("failed to update task: %w", err)
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}

	return response, nil
}

// DeleteTasksId обрабатывает запрос на удаление задачи по ID.
func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskID := request.Id // Извлекаем ID задачи из параметров запроса

	err := h.Service.DeleteTask(taskID)
	if err != nil {
		log.Printf("Error deleting task with ID %d: %v", taskID, err)
		return nil, err // Возвращаем ошибку если удаление не удалось
	}

	return tasks.DeleteTasksId204Response{}, nil // Возвращаем успешный ответ без тела
}

// NewHandler создает новый экземпляр Handler с заданным сервисом.
func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}
