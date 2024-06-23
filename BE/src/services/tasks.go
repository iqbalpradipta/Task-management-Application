package services

import (
	"github.com/iqbalpradipta/Task-management-Application/BE/src/model"
	"gorm.io/gorm"
)

type Tasks interface {
	GetTasks() ([]model.Task, error)
	GetTaskById(id int) (model.Task, error)
	CreateTask(task model.Task) (model.Task, error)
	UpdateTask(task model.Task) (model.Task, error)
	DeleteTask(task model.Task) (model.Task, error)
} 

func RepositoryTask(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetTasks() ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.Find(&tasks).Error

	return tasks, err
}

func (r *Repository) GetTaskById(id int) (model.Task, error) {
	var tasks model.Task
	err := r.db.First(&tasks, id).Error

	return tasks, err
}

func (r *Repository) CreateTask(task model.Task) (model.Task, error) {
	err := r.db.Create(&task).Error

	return task, err
}

func (r *Repository) UpdateTask(task model.Task) (model.Task, error) {
	err := r.db.Save(&task).Error

	return task, err
}

func (r *Repository) DeleteTask(task model.Task) (model.Task, error) {
	err := r.db.Delete(&task).Error

	return task, err
}