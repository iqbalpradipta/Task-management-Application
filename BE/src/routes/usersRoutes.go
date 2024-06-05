package routes

import (
	"github.com/iqbalpradipta/Task-management-Application/BE/src/config"
	"github.com/iqbalpradipta/Task-management-Application/BE/src/controllers"
	"github.com/iqbalpradipta/Task-management-Application/BE/src/middleware"
	"github.com/iqbalpradipta/Task-management-Application/BE/src/services"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Group) {
	r := services.RepositoryUsers(config.DB)
	c := controllers.UserController(r)

	e.GET("/users", c.GetUsers)
	e.GET("/users/:id", c.GetUserById, middleware.JWTMiddleware())
	e.POST("/users", c.CreateUser)
	e.POST("/users/login", c.Login)
	e.PATCH("/users/:id", c.UpdateUsers, middleware.JWTMiddleware())
	e.DELETE("/users/:id", c.DeleteUsers)
}