package main

import (
	"net/http"

	"github.com/iqbalpradipta/Task-management-Application/BE/src/config"
	"github.com/iqbalpradipta/Task-management-Application/BE/src/migration"
	"github.com/iqbalpradipta/Task-management-Application/BE/src/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	  }))

	config.DbConfig()
	migration.RunMigration()
	routes.Routes(e.Group("/api/v1"))
	
	e.Logger.Fatal(e.Start(":8000"))
}