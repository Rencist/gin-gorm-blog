package service

import (
	"context"
	"gin-gorm-blog/dto"
	"gin-gorm-blog/entity"
	"gin-gorm-blog/repository"

	"github.com/mashingan/smapping"
)

type BlogService interface {
	CreateBlog(ctx context.Context, blogDTO dto.BlogCreateDto) (entity.Blog, error)
	GetAllBlog(ctx context.Context) ([]entity.Blog, error)
	GetUserBlog(ctx context.Context, userID string) ([]entity.Blog, error)
}

type blogService struct {
	blogRepository repository.BlogRepository
}

func NewBlogService(br repository.BlogRepository) BlogService {
	return &blogService{
		blogRepository: br,
	}
}

func(bs *blogService) CreateBlog(ctx context.Context, blogDTO dto.BlogCreateDto) (entity.Blog, error) {
	blog := entity.Blog{}
	err := smapping.FillStruct(&blog, smapping.MapFields(blogDTO))
	if err != nil {
		return blog, err
	}
	return bs.blogRepository.CreateBlog(ctx, blog)
}

func(bs *blogService) GetAllBlog(ctx context.Context) ([]entity.Blog, error) {
	return bs.blogRepository.GetAllBlog(ctx)
}

func(bs *blogService) GetUserBlog(ctx context.Context, userID string) ([]entity.Blog, error) {
	return bs.blogRepository.FindBlogByUserID(ctx, userID)
}