package migration

import (
	"github.com/iqbalpradipta/Task-management-Application/BE/src/config"
	"github.com/iqbalpradipta/Task-management-Application/BE/src/model"
)

func RunMigration() {
	err := config.DB.AutoMigrate(
		&model.Users{},
		&model.Task{},
		&model.TaskDetail{},
		&model.TaskStatus{},
	)

	if err != nil {
		panic(err)
	}
}