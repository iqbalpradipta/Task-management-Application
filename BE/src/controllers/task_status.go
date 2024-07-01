package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/iqbalpradipta/Task-management-Application/BE/src/helpers"
	"github.com/iqbalpradipta/Task-management-Application/BE/src/model"
	"github.com/iqbalpradipta/Task-management-Application/BE/src/services"
	"github.com/labstack/echo/v4"
)

type taskStatusService struct {
	TaskStatusRepository services.TaskStatus
}

func TaskStatusController(taskStatusRepository services.TaskStatus) *taskStatusService {
	return &taskStatusService{taskStatusRepository}
}

func (t *taskStatusService) GetTaskStatusAll(c echo.Context) error {
	tasks, err := t.TaskStatusRepository.GetTaskStatusAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Error when get data !"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success get data Status", tasks))
}

func (t *taskStatusService) GetTaskStatusById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	tasks, err := t.TaskStatusRepository.GetTaskStatusById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Failed get data Status By Id"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success get Task Status" ,convResponsTasksStatus(tasks)))

}

func (t *taskStatusService) CreateTaskStatus(c echo.Context) error {
	taskStatusReq := new(model.TaskStatusRequest)

	if err := c.Bind(&taskStatusReq); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed when Bind data!"))
	}

	validation := validator.New()
	err := validation.Struct(taskStatusReq)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed, Please check your field !"))
	}

	taskStatus := model.TaskStatus{
		Status: taskStatusReq.Status,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	createdTaskStatus, err := t.TaskStatusRepository.CreateTaskStatus(taskStatus)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed When Create data !"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success Create data", createdTaskStatus))
}

func (t *taskStatusService) UpdateTaskStatus(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	taskStatus, err := t.TaskStatusRepository.GetTaskStatusById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed get data by id"))
 	}

	req := new(model.TaskStatusRequestUpdate)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed when bind data"))
	}
	
	if req.Status != "" {
		taskStatus.Status = req.Status
	}

	taskStatus.Updated_at = time.Now()
	response, err := t.TaskStatusRepository.UpdateTaskStatus(taskStatus)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Failed when updated data"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success Update data", response))
}

func (t *taskStatusService) DeleteTaskStatus(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	taskStatus, err := t.TaskStatusRepository.GetTaskStatusById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Failed when get data by id"))
	}

	data, err := t.TaskStatusRepository.DeleteTaskStatus(taskStatus)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed delete data"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success delete data", data))
}

func convResponsTasksStatus(t model.TaskStatus) model.TaskStatusResponse {
	return model.TaskStatusResponse{
		ID: t.ID,
		Status: t.Status,
	}
}