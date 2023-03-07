package repository

import (
	"context"
	"gin-gorm-blog/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentRepository interface {
	CreateComment(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	UpdateComment(ctx context.Context, comment entity.Comment) (error)
	ValidateCommentUser(ctx context.Context, userID string, commentID string) (bool)
	FindCommentByID(ctx context.Context, commentID string) (entity.Comment, error)
}

type commentConnection struct {
	connection *gorm.DB
	blogRepository BlogRepository
}

func NewCommentRepository(db *gorm.DB, br BlogRepository) CommentRepository {
	return &commentConnection{
		connection: db,
		blogRepository: br,
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

func(db *commentConnection) ValidateCommentUser(ctx context.Context, userID string, commentID string) (bool) {
	comment, err := db.FindCommentByID(ctx, commentID)
	if err != nil {
		return false
	}
	blog, err := db.blogRepository.CheckBlogCommentByID(ctx, comment.BlogID.String())
	if blog.UserID.String() == userID {
		return true
	}
	return false
}