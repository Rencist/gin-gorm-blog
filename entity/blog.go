package entity

import (
	"time"

	"github.com/google/uuid"
)

type Blog struct {
	ID        			uuid.UUID   `gorm:"primary_key;not_null" json:"id"`
	Title 				string 		`json:"title"`
	Slug 				string 		`json:"slug"`
	Description 		string 		`json:"description"`
	LikeCount 			int 		`json:"like_count"`
	WatchCount 			int 		`json:"watch_count"`

	UserID   			uuid.UUID 	`gorm:"foreignKey" json:"user_id"`
	User     			*User  		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
	
	Comments 			[]Comment 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments,omitempty"`
	BlogTags 			[]BlogTag 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"blog_tags,omitempty"`
	
	CreatedAt 			time.Time 	`json:"created_at" default:"CURRENT_TIMESTAMP"`
	UpdatedAt 			time.Time 	`json:"updated_at"`
}