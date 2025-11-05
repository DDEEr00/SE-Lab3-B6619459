package models

import ("gorm.io/gorm"
)

type Course struct {
	gorm.Model
	CourseCode string `json:"course_code"`
	CourseName string `json:"course_name"`
	Description string `json:"description"`
	Credit int `json:"credit"`
	Department string `json:"department"`
}