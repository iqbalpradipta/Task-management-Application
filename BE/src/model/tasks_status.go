package model

import "time"

type TaskStatus struct {
	ID          int       `json:"id" gorm:"primaryKey:autoIncrement"`
	Status		string	  `json:"status" form:"status"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

type TaskStatusRequest struct {
	Status		string	  `json:"status" form:"status" validation:"required"`
}

type TaskStatusRequestUpdate struct {
	Status		string	  `json:"status" form:"status"`
}

type TaskStatusResponse struct {
	ID          int   
	Status		string
	Created_at	time.Time
	Updated_at	time.Time
}

