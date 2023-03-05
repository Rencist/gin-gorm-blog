package repository

import (
	"context"
	"gin-gorm-blog/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(ctx context.Context, user entity.User) (entity.User, error)
	GetAllUser(ctx context.Context) ([]entity.User, error)
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func(db *userConnection) RegisterUser(ctx context.Context, user entity.User) (entity.User, error) {
	user.ID = uuid.New()
	uc := db.connection.Create(&user)
	if uc.Error != nil {
		return entity.User{}, uc.Error
	}
	return user, nil
}

func(db *userConnection) GetAllUser(ctx context.Context) ([]entity.User, error) {
	var listUser []entity.User
	tx := db.connection.Find(&listUser)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listUser, nil
}