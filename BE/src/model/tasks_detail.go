package model

import "time"

type TaskDetail struct {
	ID          int       `json:"id" gorm:"primaryKey:autoIncrement"`
	Archived 	bool      `json:"archived" form:"archived"`
	Priority	string	  `json:"priority" form:"priority"`
	Tag			string	  `json:"tag" form:"tag"`
	DueDate		time.Time `json:"duedate" form:"duedate"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

type TaskDetailRequest struct {
	Archived 	bool      `json:"archived" form:"archived" validation:"required"`
	Priority	string	  `json:"priority" form:"priority" validation:"required"`
	Tag			string	  `json:"tag" form:"tag" validation:"required"`
	DueDate		time.Time `json:"duedate" form:"duedate" validation:"required"`
}

type TaskDetailRequestUpdate struct {
	Archived 	bool      `json:"archived" form:"archived"`
	Priority	string	  `json:"priority" form:"priority"`
	Tag			string	  `json:"tag" form:"tag"`
	DueDate		time.Time `json:"duedate" form:"duedate"`
}

type TasksDetailResponse struct {
	ID          int       
	Archived 	bool     
	Priority	string	 
	Tag			string
	DueDate		time.Time
}