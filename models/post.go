package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title       string
	Description string
	AuthorID    uint
	Author      User `gorm:"foreignKey:AuthorID"`
	PublishedAt time.Time
	IsPublished bool `gorm:"default:false"`
}
