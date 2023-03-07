package service

import (
	"context"
	"gin-gorm-blog/dto"
	"gin-gorm-blog/entity"
	"gin-gorm-blog/repository"

	"github.com/mashingan/smapping"
)

type CommentService interface {
	CreateComment(ctx context.Context, commentDTO dto.CommentCreateDto) (entity.Comment, error)
	UpdateComment(ctx context.Context, commentDTO dto.CommentCreateDto) (error)
	ValidateCommentUser(ctx context.Context, userID string, commentID string) (bool)
}

type commentService struct {
	commentRepository repository.CommentRepository
	blogRepository repository.BlogRepository
}

func NewCommentService(cc repository.CommentRepository, br repository.BlogRepository) CommentService {
	return &commentService{
		commentRepository: cc,
		blogRepository: br,
	}
}

func(cs *commentService) CreateComment(ctx context.Context, commentDTO dto.CommentCreateDto) (entity.Comment, error) {
	comment := entity.Comment{}
	err := smapping.FillStruct(&comment, smapping.MapFields(commentDTO))
	if err != nil {
		return comment, err
	}
	return cs.commentRepository.CreateComment(ctx, comment)
}

func(cs *commentService) UpdateComment(ctx context.Context, commentDTO dto.CommentCreateDto) (error) {
	comment := entity.Comment{}
	err := smapping.FillStruct(&comment, smapping.MapFields(commentDTO))
	if err != nil {
		return err
	}
	return cs.commentRepository.UpdateComment(ctx, comment)
}

func(cs *commentService) ValidateCommentUser(ctx context.Context, userID string, commentID string) (bool) {
	comment, err := cs.commentRepository.FindCommentByID(ctx, commentID)
	if err != nil {
		return false
	}
	blog, err := cs.blogRepository.CheckBlogCommentByID(ctx, comment.BlogID.String())
	if blog.UserID.String() == userID {
		return true
	}
	return false
}