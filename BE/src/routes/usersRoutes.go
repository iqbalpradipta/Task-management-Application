package routes

import (
	"github.com/iqbalpradipta/Task-management-Application/BE/src/config"
	"github.com/iqbalpradipta/Task-management-Application/BE/src/controllers"
	"github.com/iqbalpradipta/Task-management-Application/BE/src/services"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Group) {
	r := services.RepositoryUsers(config.DB)
	c := controllers.UserController(r)

	e.GET("/users", c.GetUsers)
	e.GET("/users/:id", c.GetUserById)
	e.POST("/users", c.CreateUser)
	e.PATCH("/users/:id", c.UpdateUsers)
	e.DELETE("/users/:id", c.DeleteUsers)
}