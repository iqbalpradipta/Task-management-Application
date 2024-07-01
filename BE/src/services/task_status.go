package services

import (
	"github.com/iqbalpradipta/Task-management-Application/BE/src/model"
	"gorm.io/gorm"
)

type TaskStatus interface {
	GetTaskStatusAll() ([]model.TaskStatus, error)
	GetTaskStatusById(id int)(model.TaskStatus, error)
	CreateTaskStatus(taskStatus model.TaskStatus) (model.TaskStatus, error)
	UpdateTaskStatus(taskStatus model.TaskStatus) (model.TaskStatus, error)
	DeleteTaskStatus(taskStatus model.TaskStatus) (model.TaskStatus, error)
}

func TaskStatusRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository)GetTaskStatusAll() ([]model.TaskStatus, error) {
	var taskStatus []model.TaskStatus

	err := r.db.Find(&taskStatus).Error

	return taskStatus, err
}

func (r *Repository) GetTaskStatusById(id int) (model.TaskStatus, error) {
	var taskStatus model.TaskStatus

	err := r.db.First(&taskStatus, id).Error

	return taskStatus, err
}

func (r *Repository) CreateTaskStatus(taskStatus model.TaskStatus) (model.TaskStatus, error) {
	err := r.db.Create(&taskStatus).Error

	return taskStatus, err
}

func (r *Repository) UpdateTaskStatus(taskStatus model.TaskStatus) (model.TaskStatus, error)  {
	err := r.db.Save(&taskStatus).Error

	return taskStatus, err
}

func (r *Repository) DeleteTaskStatus(taskStatus model.TaskStatus) (model.TaskStatus, error) {
	err := r.db.Delete(&taskStatus).Error

	return taskStatus, err 
}