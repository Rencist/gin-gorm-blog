package entity

import (
	"github.com/google/uuid"
)

type Tag struct {
	ID        			uuid.UUID   `gorm:"primary_key;not_null" json:"id"`
	Name 				string 		`json:"name"`

	BlogTags 			[]BlogTag 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"blog_tags,omitempty"`
	
	Timestamp
}