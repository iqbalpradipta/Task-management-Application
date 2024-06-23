package services

import (
	"github.com/iqbalpradipta/Task-management-Application/BE/src/model"
	"gorm.io/gorm"
)

type Users interface {
	GetUsers() ([]model.Users , error)
	GetUserById(id int) (model.Users, error)
	GetUserByEmail(email string) (model.Users, error)
	CreateUsers(users model.Users) (model.Users, error)
	Login(users model.Users) (model.Users, error)
	UpdatedUsers(users model.Users) (model.Users, error)
	DeleteUsers(users model.Users) (model.Users, error)
}

func RepositoryUsers(db *gorm.DB) *Repository{
	return &Repository{db}
}

 
func (r *Repository) GetUsers() ([]model.Users, error) {
	var users []model.Users
	err := r.db.Find(&users).Error

	return users, err
}

func (r *Repository) GetUserById(id int) (model.Users, error) {
	var users model.Users
	err := r.db.First(&users, id).Error

	return users, err
}

func (r *Repository) GetUserByEmail(email string) (model.Users, error) {
	var users model.Users
	err := r.db.Where("email = ?", email).First(&users).Error

	return users, err
}

func (r *Repository) CreateUsers(users model.Users) (model.Users, error) {
	err := r.db.Create(&users).Error

	return users, err
}

func (r *Repository) Login(users model.Users) (model.Users, error) {
	err := r.db.First(&users).Error

	return users, err
}

func (r *Repository) UpdatedUsers(users model.Users) (model.Users, error) {
	err := r.db.Save(&users).Error

	return users, err
}

func (r *Repository) DeleteUsers(users model.Users) (model.Users, error) {
	err := r.db.Delete(&users).Error

	return users, err
}
