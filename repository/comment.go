package repository

import (
	"context"
	"gin-gorm-blog/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentRepository interface {
	CreateComment(ctx context.Context, comment entity.Comment) (entity.Comment, error)
}

type commentConnection struct {
	connection *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentConnection{
		connection: db,
	}
}

func(db *commentConnection) CreateComment(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	comment.ID = uuid.New()
	cc := db.connection.Create(&comment)
	if cc.Error != nil {
		return entity.Comment{}, cc.Error
	}
	return comment, nil
}