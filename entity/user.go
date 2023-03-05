package entity

import (
	"gin-gorm-blog/helpers"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	u.Password, err = helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}