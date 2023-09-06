package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"Username" gorm:"unique"`
	Password string `json:"Password"`
	Role     string `json:"Role" gorm:"default:admin"`
}
