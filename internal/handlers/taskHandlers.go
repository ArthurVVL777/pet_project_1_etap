package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"pet_project_1_etap/internal/taskService"
	"pet_project_1_etap/internal/web/tasks"
)

type Handler struct {
	Service *taskService.TaskService
}

// NewHandler создает новый экземпляр Handler с заданным сервисом.
func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) GetTasks(ctx echo.Context) error {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		log.Printf("Error fetching tasks: %v", err)
		return ctx.JSON(http.StatusInternalServerError, "Error fetching tasks")
	}

	// Формируем список задач с добавлением поля user_id
	response := make(tasks.GetTasks200JSONResponse, 0, len(allTasks))
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
			UserID: &tsk.UserID,
		}
		response = append(response, task)
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) PostTasks(ctx echo.Context) error {
	var request tasks.PostTaskRequestBody
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request body")
	}

	if request.Task == nil || *request.Task == "" || request.UserID == nil {
		return ctx.JSON(http.StatusBadRequest, "Task and UserID must not be empty")
	}

	createdTask, err := h.Service.CreateTask(taskService.Task{
		Task:   *request.Task,
		IsDone: *request.IsDone,
		UserID: *request.UserID,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error creating task")
	}

	response := tasks.Task{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	return ctx.JSON(http.StatusCreated, response)
}

func (h *Handler) PatchTasksId(ctx echo.Context, id uint) error {
	var request tasks.PatchTaskRequestBody
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request body")
	}

	existingTask, err := h.Service.GetTaskByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "Task not found")
	}

	// Обновляем только те поля, которые были указаны в запросе
	if request.Task != nil && *request.Task != "" {
		existingTask.Task = *request.Task
	}
	if request.IsDone != nil {
		existingTask.IsDone = *request.IsDone
	}

	updatedTask, err := h.Service.UpdateTaskByID(id, existingTask)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error updating task")
	}

	response := tasks.Task{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) DeleteTasksId(ctx echo.Context, id uint) error {
	err := h.Service.DeleteTask(id)
	if err != nil {
		log.Printf("Error deleting task with ID %d: %v", id, err)
		return ctx.JSON(http.StatusInternalServerError, "Error deleting task")
	}

	return ctx.NoContent(http.StatusNoContent) // Возвращаем успешный ответ без тела
}
