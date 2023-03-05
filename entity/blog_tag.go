package entity

import (
	"github.com/google/uuid"
)

type BlogTag struct {
	ID               uuid.UUID  `gorm:"primary_key;not_null" json:"id"`

	BlogID   		uuid.UUID 	`gorm:"foreignKey" json:"blog_id"`
	Blog     		*Blog  		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"blog,omitempty"`
	TagID   		uuid.UUID 	`gorm:"foreignKey" json:"tag_id"`
	Tag     		*Tag  		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"tag,omitempty"`

	Timestamp
}