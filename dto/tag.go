package dto

import (
	"github.com/google/uuid"
)

type TagCreateDto struct {
	ID        			uuid.UUID   `gorm:"primary_key" json:"id" form:"id"`
	Name 				string 		`json:"name" form:"name"`
}