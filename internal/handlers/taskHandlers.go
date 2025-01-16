package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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

	response := make(tasks.GetTasks200JSONResponse, 0, len(allTasks))
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
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

	return ctx.JSON(http.StatusCreated, createdTask)
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

	return ctx.JSON(http.StatusOK, updatedTask)
}

func (h *Handler) DeleteTasksId(ctx echo.Context, id uint) error {
	err := h.Service.DeleteTask(id)
	if err != nil {
		log.Printf("Error deleting task with ID %d: %v", id, err)
		return ctx.JSON(http.StatusInternalServerError, "Error deleting task")
	}

	return ctx.NoContent(http.StatusNoContent) // Возвращаем успешный ответ без тела
}

func (h *UserHandler) GetTasksForUser(ctx echo.Context) error {
	userID := ctx.Param("user_id") // Получаем user_id из URL

	// Преобразуем строку в uint
	id, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	tasks, err := h.Service.GetTasksForUser(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error fetching tasks")
	}

	return ctx.JSON(http.StatusOK, tasks)
}
