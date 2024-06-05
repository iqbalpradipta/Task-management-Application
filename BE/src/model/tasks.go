package model

import "time"

type Task struct {
	ID          		int      `json:"id" gorm:"primaryKey:autoIncrement"`
	Description 		string    `json:"description" form:"description"`
	UserID				int
	Task_Detail_Id		int
	Task_Detail			TaskDetail	`gorm:"foreignKey:Task_Detail_Id"`
	Task_Status_Id		int
	Task_Status			TaskStatus	`gorm:"foreignKey:Task_Status_Id"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

type TaskRequest struct {
	Description	string	  	`json:"description" form:"description" validation:"required"`
}

type TaskRequestUpdate struct {
	Description string 		`json:"description" form:"description"`
}

type TaskResponse struct {
	ID			int
	Description	string
	Created_at	time.Time
	Updated_at	time.Time
}

