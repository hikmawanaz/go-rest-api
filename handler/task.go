package handler

import (
	"belyfe/auth"
	"belyfe/helper"
	"belyfe/user"
	"net/http"

	"belyfe/task"

	"github.com/gin-gonic/gin"
)

type taskHandler struct {
	taskService task.Service
	authService auth.Service
}

func NewTaskHandler(taskService task.Service, authService auth.Service) *taskHandler {
	return &taskHandler{taskService, authService}
}

func (h *taskHandler) CreateTask(c *gin.Context) {
	var input task.Task

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("create task failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	input.UserID = uint(userID)

	newTask, err := h.taskService.CreateTask(input)

	if err != nil {
		response := helper.APIResponse("create task failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("new task created", http.StatusOK, "success", newTask)

	c.JSON(http.StatusOK, response)
}

func (h *taskHandler) FindAllTasks(c *gin.Context) {
	tasks, err := h.taskService.GetAllTasks()

	if err != nil {
		response := helper.APIResponse("failed to get tasks", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success get all tasks", http.StatusOK, "success", tasks)

	c.JSON(http.StatusOK, response)
}

func (h *taskHandler) CompleteSingleTask(c *gin.Context) {
	var input task.CompleteSingleTaskInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("failed to complete task", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.taskService.CompleteSingleTask(input)

	if err != nil {
		response := helper.APIResponse("create task failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	tasks, err := h.taskService.GetAllTasks()

	if err != nil {
		response := helper.APIResponse("failed to get tasks", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success get all tasks", http.StatusOK, "success", tasks)

	c.JSON(http.StatusOK, response)
}
