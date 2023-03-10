package repository

import (
	"context"
	"gin-gorm-blog/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentRepository interface {
	GetTotalDataByBlogID(ctx context.Context, blogID string) (int64, error)
	CreateComment(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	UpdateComment(ctx context.Context, comment entity.Comment) (error)
	FindCommentByID(ctx context.Context, commentID string) (entity.Comment, error)
}

type commentConnection struct {
	connection *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentConnection{
		connection: db,
	}
}

func(db *commentConnection) GetTotalDataByBlogID(ctx context.Context, blogID string) (int64, error) {
	var totalData int64
	cc := db.connection.Model(&entity.Comment{}).Where("blog_id = ?", blogID).Count(&totalData)
	if cc.Error != nil {
		return 0, cc.Error
	}
	return totalData, nil
}

func(db *commentConnection) CreateComment(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	comment.ID = uuid.New()
	cc := db.connection.Create(&comment)
	if cc.Error != nil {
		return entity.Comment{}, cc.Error
	}
	return comment, nil
}

func(db *commentConnection) UpdateComment(ctx context.Context, comment entity.Comment) (error) {
	bc := db.connection.Updates(&comment)
	if bc.Error != nil {
		return bc.Error
	}
	return nil
}

func(db *commentConnection) FindCommentByID(ctx context.Context, commentID string) (entity.Comment, error) {
	var comment entity.Comment
	bc := db.connection.Where("id = ?", commentID).Find(&comment)
	if bc.Error != nil {
		return entity.Comment{}, nil
	}
	return comment, nil
}