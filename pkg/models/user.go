package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID            uint   `json:"id" gorm:"primaryKey"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Preferences   string `json:"preferences"`
	Accessibility string `json:"accessibility"`
}
