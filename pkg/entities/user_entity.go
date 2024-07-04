package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"not null;unique" json:"username"`
	Email    string `gorm:"not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
