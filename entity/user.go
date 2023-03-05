package entity

import (
	"github.com/google/uuid"
)

type User struct {
	ID        	uuid.UUID   `gorm:"primary_key;not_null" json:"id"`
	Name 		string 		`json:"name"`
	Email 		string 		`json:"email" binding:"email"`
	NoTelp 		string 		`json:"no_telp"`
	Password 	string  	`json:"password"`
	
	Blogs 		[]Blog 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"blogs,omitempty"`
	
	Timestamp
}