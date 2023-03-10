package dto

import (
	"github.com/google/uuid"
)

type BlogTagCreateDto struct {
	ID        		uuid.UUID   `gorm:"primary_key" json:"id" form:"id"`

	BlogID 		string 	`gorm:"foreignKey" json:"blog_id" form:"blog_id" binding:"required"`
	TagID 		string 	`gorm:"foreignKey" json:"tag_id" form:"tag_id" binding:"required"`
}