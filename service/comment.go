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
}

func NewCommentService(cc repository.CommentRepository) CommentService {
	return &commentService{
		commentRepository: cc,
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
	return cs.commentRepository.ValidateCommentUser(ctx, userID, commentID) 
}