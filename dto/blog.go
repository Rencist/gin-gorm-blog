package dto

import (
	"gin-gorm-blog/entity"

	"github.com/google/uuid"
)

type BlogCreateDto struct {
	ID        		uuid.UUID   `gorm:"primary_key" json:"id" form:"id"`
	Title 			string 		`json:"title" form:"title" binding:"required"`
	Slug 			string 		`json:"slug" form:"slug" binding:"required"`
	Description 	string 		`json:"description" form:"description" binding:"required"`

	Tag 			TagCreateDto `json:"tag" form:"tag"`

	UserID 			uuid.UUID 	`gorm:"foreignKey" json:"user_id" form:"user_id"`
}

type BlogUpdateDto struct {
	ID        		uuid.UUID   `gorm:"primary_key" json:"id" form:"id"`
	Title 			string 		`json:"title" form:"title"`
	Slug 			string 		`json:"slug" form:"slug"`
	Description 	string 		`json:"description" form:"description"`

	UserID 		uuid.UUID 	`gorm:"foreignKey" json:"user_id" form:"user_id"`
}

type BlogResponseDto struct {
	ID        			uuid.UUID   `gorm:"primary_key;not_null" json:"id"`
	UserName			string		`json:"user_name"`
	Title 				string 		`json:"title"`
	Slug 				string 		`json:"slug"`
	Description 		string 		`json:"description"`
	LikeCount 			int 		`json:"like_count"`
	WatchCount 			int 		`json:"watch_count"`

	UserID   			uuid.UUID 	`gorm:"foreignKey" json:"user_id"`
	User     			*entity.User  		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
	
	Comments 			[]entity.Comment 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments,omitempty"`
	BlogTags 			[]entity.BlogTag 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"blog_tags,omitempty"`
	
	entity.Timestamp
}