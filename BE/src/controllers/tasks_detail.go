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

type taskDetailService struct {
	TaskDetailRepository services.TasksDetail
}

func TaskDetailController(taskDetailRepository services.TasksDetail) *taskDetailService {
	return &taskDetailService{taskDetailRepository}
}

func (t *taskDetailService) GetTaskDetailAll(c echo.Context) error {
	tasks, err := t.TaskDetailRepository.GetTaskDetailAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Error when get data !"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success get data detail", tasks))
}

func (t *taskDetailService) GetTaskDetailById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	tasks, err := t.TaskDetailRepository.GetTaskDetailById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Failed get data detail By Id"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success get Task Detail" ,convResponsTasksDetail(tasks)))

}

func (t *taskDetailService) CreateTaskDetail(c echo.Context) error {
	taskDetailReq := new(model.TaskDetailRequest)

	if err := c.Bind(&taskDetailReq); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed when Bind data!"))
	}

	validation := validator.New()
	err := validation.Struct(taskDetailReq)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed, Please check your field !"))
	}

	taskDetail := model.TaskDetail{
		Archived: taskDetailReq.Archived,
		Priority: taskDetailReq.Priority,
		Tag: taskDetailReq.Tag,
		DueDate: taskDetailReq.DueDate,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	createdTaskDetail, err := t.TaskDetailRepository.CreateTaskDetail(taskDetail)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed When Create data !"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success Create data", createdTaskDetail))
}

func (t *taskDetailService) UpdateTaskDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	taskDetail, err := t.TaskDetailRepository.GetTaskDetailById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed get data by id"))
 	}

	req := new(model.TaskDetailRequestUpdate)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed when bind data"))
	}
	
	if req.Tag != "" {
		taskDetail.Tag = req.Tag
	}

	if req.Priority != "" {
		taskDetail.Priority = req.Priority
	}

	taskDetail.Updated_at = time.Now()
	response, err := t.TaskDetailRepository.UpdateTaskDetail(taskDetail)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Failed when updated data"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success Update data", response))
}

func (t *taskDetailService) DeleteTaskDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	taskDetail, err := t.TaskDetailRepository.GetTaskDetailById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Failed when get data by id"))
	}

	data, err := t.TaskDetailRepository.DeleteTaskDetail(taskDetail)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed delete data"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success delete data", data))
}

func convResponsTasksDetail(t model.TaskDetail) model.TasksDetailResponse {
	return model.TasksDetailResponse{
		ID: t.ID,
		Archived: t.Archived,
		Priority: t.Priority,
		Tag: t.Tag,
		DueDate: t.DueDate,
	}
}