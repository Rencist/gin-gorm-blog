package dto

import (
	"github.com/google/uuid"
)

type BlogCreateDto struct {
	ID        		uuid.UUID   `gorm:"primary_key" json:"id" form:"id"`
	Title 			string 		`json:"title" form:"title"`
	Slug 			string 		`json:"slug" form:"slug"`
	Description 	string 		`json:"description" form:"description"`
	LikeCount 		int 		`json:"like_count" form:"like_count"`
	WatchCount 		int 		`json:"watch_count" form:"watch_count"`

	UserID 		uuid.UUID 	`gorm:"foreignKey" json:"user_id" form:"user_id" binding:"required"`
}

type BlogUpdateDto struct {
	ID        		uuid.UUID   `gorm:"primary_key" json:"id" form:"id"`
	Title 			string 		`json:"title" form:"title" binding:"required"`
	Slug 			string 		`json:"slug" form:"slug" binding:"required"`
	Description 	string 		`json:"description" form:"description" binding:"required"`
	LikeCount 		int 		`json:"like_count" form:"like_count" binding:"required"`
	WatchCount 		int 		`json:"watch_count" form:"watch_count" binding:"required"`

	UserID 		uuid.UUID 	`gorm:"foreignKey" json:"user_id" form:"user_id" binding:"required"`
}