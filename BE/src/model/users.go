package model

import "time"

type Users struct {
	ID         int       	`json:"id" gorm:"primaryKey:autoIncrement"`
	Name       string    	`json:"name"`
	Email      string    	`json:"email"`
	Password   string    	`json:"password"`
	TasksID	   []Task		`gorm:"foreignKey:UserID"`
	Created_at time.Time	`json:"created_at"`
	Updated_at	time.Time	`json:"updated_at"`
}

type UsersRequest struct {
	Name	string `json:"name" form:"name" validation:"required"`
	Email	string `json:"email" form:"email" validation:"required"`
	Password string `json:"password" form:"password" validation:"required"`
}

type UsersRequestUpdate struct {
	Name	string `json:"name" form:"name"`
	Email	string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserResponse struct {
	ID		int		`json:"id"`
	Name	string  `json:"name"`
	Email	string	`json:"email"`
	Created_at time.Time	`json:"created_at`
	Updated_at time.Time	`json:"updated_at`
}

type Login struct {
	Email	string	`json:"email"`
	Password string	`json:"password"`
}