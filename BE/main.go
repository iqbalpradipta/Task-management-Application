package main

import (
	"github.com/iqbalpradipta/Task-management-Application/src/config"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.DbConfig()
	// migration.RunMigration()
	// routes.Routes(e.Group("/api/v1"))
	
	e.Logger.Fatal(e.Start(":8000"))
}