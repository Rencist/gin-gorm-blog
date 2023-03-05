package dto

import (
	"github.com/google/uuid"
)

type BlogTagCreateDto struct {
	ID        		uuid.UUID   `gorm:"primary_key" json:"id" form:"id"`

	BlogID 		uuid.UUID 	`gorm:"foreignKey" json:"blog_id" form:"blog_id" binding:"required"`
	TagID 		uuid.UUID 	`gorm:"foreignKey" json:"tag_id" form:"tag_id" binding:"required"`
}