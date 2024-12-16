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

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}
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

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body
	if taskRequest == nil || taskRequest.Task == nil || taskRequest.IsDone == nil {
		return nil, fmt.Errorf("request body or fields cannot be nil")
	}

	taskToCreate := taskService.Task{
		Text:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}

	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Text,
		IsDone: &createdTask.IsDone,
	}

	return response, nil
}

func (h *Handler) PatchTasks(ctx context.Context, request tasks.PatchTasksRequestObject) (tasks.PatchTasksResponseObject, error) {
	taskID := *request.Body.Id // Разыменовываем указатель на ID задачи
	taskRequest := request.Body

	if taskRequest == nil {
		return nil, fmt.Errorf("request body cannot be nil")
	}

	// Получаем существующую задачу из базы данных
	existingTask, err := h.Service.GetTaskByID(taskID)
	if err != nil {
		return nil, fmt.Errorf("task not found: %v", err)
	}

	// Обновляем только те поля, которые были переданы в запросе
	if taskRequest.Task != nil {
		existingTask.Task = *taskRequest.Task // Обновляем текст задачи только если он не nil
	}

	if taskRequest.IsDone != nil {
		existingTask.IsDone = *taskRequest.IsDone // Обновляем статус завершенности только если он не nil
	}

	// Сохраняем обновленную задачу в базе данных
	updatedTask, err := h.Service.UpdateTaskByID(existingTask)
	if err != nil {
		log.Printf("Error updating task with ID %d: %v", taskID, err)
		return nil, err
	}

	response := tasks.PatchTasks200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}

	return response, nil
}

func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskID := request.Id // Извлекаем ID задачи из параметров запроса

	err := h.Service.DeleteTask(taskID)
	if err != nil {
		return nil, err // Возвращаем ошибку, если удаление не удалось
	}

	return tasks.DeleteTasksId204Response{}, nil // Возвращаем успешный ответ без тела
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}
