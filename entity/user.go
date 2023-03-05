package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        	uuid.UUID   `gorm:"primary_key;not_null" json:"id"`
	Name 		string 		`json:"name"`
	Email 		string 		`json:"email" binding:"email"`
	NoTelp 		string 		`json:"no_telp"`
	Password 	string  	`json:"password"`
	
	Blogs 		[]Blog 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"blogs,omitempty"`
	
	CreatedAt 	time.Time 	`json:"created_at" default:"CURRENT_TIMESTAMP"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
}