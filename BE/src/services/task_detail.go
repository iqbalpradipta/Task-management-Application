package services

import (
	"github.com/iqbalpradipta/Task-management-Application/BE/src/model"
	"gorm.io/gorm"
)

type TasksDetail interface {
	GetTaskDetailAll() ([]model.TaskDetail, error)
	GetTaskDetailById(id int) (model.TaskDetail, error)
	CreateTaskDetail(task model.TaskDetail) (model.TaskDetail, error)
	UpdateTaskDetail(task model.TaskDetail) (model.TaskDetail, error)
	DeleteTaskDetail(task model.TaskDetail) (model.TaskDetail, error)
}

func RepositoryTaskDetail(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetTaskDetailAll() ([]model.TaskDetail, error) {
	var taskDetails []model.TaskDetail

	err := r.db.Find(&taskDetails).Error

	return taskDetails, err
}

func (r *Repository) GetTaskDetailById(id int) (model.TaskDetail, error) {
	var taskDetail model.TaskDetail

	err := r.db.First(&taskDetail, id).Error

	return taskDetail, err
}

func (r *Repository) CreateTaskDetail(taskDetail model.TaskDetail) (model.TaskDetail, error) {
	err := r.db.Create(&taskDetail).Error

	return taskDetail, err
}

func (r *Repository) UpdateTaskDetail(taskDetail model.TaskDetail) (model.TaskDetail, error) {
	err := r.db.Save(&taskDetail).Error

	return taskDetail, err
}

func (r *Repository) DeleteTaskDetail(taskDetail model.TaskDetail) (model.TaskDetail, error) {
	err := r.db.Delete(&taskDetail).Error

	return taskDetail, err
}