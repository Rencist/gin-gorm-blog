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
	GetBlogByID(ctx context.Context, pagination entity.Pagination, blogID string) (dto.BlogPaginationResponse, error)
	LikeBlogByID(ctx context.Context, blogID string) (error)
	ValidateBlogUser(ctx context.Context, userID string, blogID string) (bool)
	UpdateBlog(ctx context.Context, blogDTO dto.BlogUpdateDto) (error)
	GetAllBlogPagination(ctx context.Context, pagination entity.Pagination) (dto.PaginationResponse, error)
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

func(bs *blogService) GetBlogByID(ctx context.Context, pagination entity.Pagination, blogID string) (dto.BlogPaginationResponse, error) {
	return bs.blogRepository.FindBlogByID(ctx, pagination, blogID)
}

func(bs *blogService) LikeBlogByID(ctx context.Context, blogID string) (error) {
	return bs.blogRepository.LikeBlogByID(ctx, blogID)
}

func(bs *blogService) ValidateBlogUser(ctx context.Context, userID string, blogID string) (bool) {
	blog, err := bs.blogRepository.CheckBlogCommentByID(ctx, blogID)
	if err != nil {
		return false
	}
	if blog.UserID.String() == userID {
		return true
	}
	return false
}

func(bs *blogService) UpdateBlog(ctx context.Context, blogDTO dto.BlogUpdateDto) (error) {
	blog := entity.Blog{}
	err := smapping.FillStruct(&blog, smapping.MapFields(blogDTO))
	if err != nil {
		return err
	}
	return bs.blogRepository.UpdateBlog(ctx, blog)
}

func(bs *blogService) GetAllBlogPagination(ctx context.Context, pagination entity.Pagination) (dto.PaginationResponse, error) {
	return bs.blogRepository.GetAllBlogPagination(ctx, pagination)
}