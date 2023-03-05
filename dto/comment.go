package dto

import (
	"github.com/google/uuid"
)

type CommentCreateDto struct {
	ID        			uuid.UUID   `gorm:"primary_key" json:"id" form:"id"`
	Description 		string 		`json:"description" form:"description"`

	BlogID 		uuid.UUID 	`gorm:"foreignKey" json:"blog_id" form:"blog_id" binding:"required"`
}