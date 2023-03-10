package entity

import (
	"github.com/google/uuid"
)

type Comment struct {
	ID        			uuid.UUID   `gorm:"primary_key;not_null" json:"id"`
	Description 		string 		`json:"description"`

	BlogID   			uuid.UUID 	`gorm:"foreignKey" json:"blog_id"`
	Blog     			*Blog  		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"blog,omitempty"`
	
	Timestamp
}