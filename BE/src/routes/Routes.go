package routes

import (
	"github.com/iqbalpradipta/Task-management-Application/BE/src/config"
	"github.com/iqbalpradipta/Task-management-Application/BE/src/controllers"
	"github.com/iqbalpradipta/Task-management-Application/BE/src/middleware"
	"github.com/iqbalpradipta/Task-management-Application/BE/src/services"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Group) {

	rusers := services.RepositoryUsers(config.DB)
	cusers := controllers.UserController(rusers)

	rtask := services.RepositoryTask(config.DB)
	ctask := controllers.TaskController(rtask)

	rtaskDetail := services.RepositoryTaskDetail(config.DB)
	cTaskDetail := controllers.TaskDetailController(rtaskDetail)

	rtaskStatus:= services.TaskStatusRepository(config.DB)
	cTaskStatus:= controllers.TaskStatusController(rtaskStatus)


	e.GET("/users", cusers.GetUsers)
	e.GET("/users/:id", cusers.GetUserById, middleware.JWTMiddleware())
	e.POST("/users", cusers.CreateUser)
	e.POST("/users/login", cusers.Login)
	e.PATCH("/users/:id", cusers.UpdateUsers, middleware.JWTMiddleware())
	e.DELETE("/users/:id", cusers.DeleteUsers)

	e.GET("/task", ctask.GetTasks)
	e.GET("/task/:id", ctask.GetTaskById)
	e.POST("/task", ctask.CreateTask, middleware.JWTMiddleware())

	e.GET("/taskDetail", cTaskDetail.GetTaskDetailAll)
	e.GET("/taskDetail/:id", cTaskDetail.GetTaskDetailById)
	e.POST("/taskDetail", cTaskDetail.CreateTaskDetail)
	e.PATCH("/taskDetail/:id", cTaskDetail.UpdateTaskDetail)
	e.DELETE("/taskDetail/:id", cTaskDetail.DeleteTaskDetail)

	e.GET("/taskStatus", cTaskStatus.GetTaskStatusAll)
	e.GET("/taskStatus/:id", cTaskStatus.GetTaskStatusById)
	e.POST("/taskStatus", cTaskStatus.CreateTaskStatus)
	e.PATCH("/taskStatus/:id", cTaskStatus.UpdateTaskStatus)
	e.DELETE("/taskStatus/:id", cTaskStatus.DeleteTaskStatus)
}