package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/iqbalpradipta/Task-management-Application/BE/src/helpers"
	"github.com/iqbalpradipta/Task-management-Application/BE/src/middleware"
	"github.com/iqbalpradipta/Task-management-Application/BE/src/model"
	"github.com/iqbalpradipta/Task-management-Application/BE/src/services"
	"github.com/labstack/echo/v4"
)

type taskService struct {
	TaskRepository services.Tasks
}

func TaskController(taskRepository services.Tasks) *taskService {
	return &taskService{taskRepository}
}

func (t *taskService) GetTasks(c echo.Context) error {
	tasks, err := t.TaskRepository.GetTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Error when get Tasks"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success get Tasks", tasks))
}

func (t *taskService) GetTaskById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	tasks, err := t.TaskRepository.GetTaskById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Error when get Tasks with Id"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success get Tasks", convResponseTasks(tasks)))
}

func (t *taskService) CreateTask(c echo.Context) error {
	taskReq := new(model.TaskRequest)

	if err := c.Bind(&taskReq); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed when bind data"))
	}

	validation := validator.New()
	err := validation.Struct(taskReq)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Your data not valid !"))
	}

	token := middleware.ExtractToken(c)

	task := model.Task{
		Description: taskReq.Description,
		UserID: token,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	createdTask, err := t.TaskRepository.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed create task"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success create Task", createdTask))
}

func convResponseTasks(t model.Task) model.TaskResponse {
	return model.TaskResponse{
		ID: t.ID,
		Description: t.Description,
		Created_at: t.Created_at,
		Updated_at: t.Updated_at,
	}
}