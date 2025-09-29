package model

import "time"

type Student struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" binding:"required,min=2,max=100"`
	DateOfBirth time.Time `json:"date_of_birth" binding:"required"`
	StudyYear   int       `json:"study_year" binding:"required,min=1,max=6"`
}

type Pagination struct {
	Limit  int `form:"limit" binding:"required,min=1,max=10"`
	Offset int `form:"offset" binding:"min=0"`
}

type UpdateStudent struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Name        *string    `json:"name" binding:"omitempty,min=2,max=100"`
	DateOfBirth *time.Time `json:"date_of_birth" binding:"omitempty"`
	StudyYear   *int       `json:"study_year" binding:"omitempty,min=1,max=6"`
}
