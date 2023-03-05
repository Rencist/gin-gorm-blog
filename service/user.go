package service

import (
	"context"
	"gin-gorm-blog/dto"
	"gin-gorm-blog/entity"
	"gin-gorm-blog/repository"

	"github.com/mashingan/smapping"
)

type UserService interface {
	RegisterUser(ctx context.Context, userDTO dto.UserCreateDto) (entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{
		userRepository: ur,
	}
}

func(us *userService) RegisterUser(ctx context.Context, userDTO dto.UserCreateDto) (entity.User, error) {
	user := entity.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(userDTO))
	user.Role = "user"
	if err != nil {
		return user, err
	}
	return us.userRepository.RegisterUser(ctx, user)
}