package service

import (
	"context"
	"gin-gorm-blog/dto"
	"gin-gorm-blog/entity"
	"gin-gorm-blog/repository"

	"github.com/google/uuid"
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
	AssignTag(ctx context.Context, blogTagDTO dto.BlogTagCreateDto) (error)
}

type blogService struct {
	blogRepository repository.BlogRepository
	tagRepository repository.TagRepository
	blogTagRepository repository.BlogTagRepository
}

func NewBlogService(br repository.BlogRepository, tr repository.TagRepository, btr repository.BlogTagRepository) BlogService {
	return &blogService{
		blogRepository: br,
		tagRepository: tr,
		blogTagRepository: btr,
	}
}

func(bs *blogService) CreateBlog(ctx context.Context, blogDTO dto.BlogCreateDto) (entity.Blog, error) {
	blog := entity.Blog{}
	err := smapping.FillStruct(&blog, smapping.MapFields(blogDTO))
	if err != nil {
		return blog, err
	}
	blogCreate, err := bs.blogRepository.CreateBlog(ctx, blog)
	if blogDTO.Tag.Name != "" {
		tag, err := bs.tagRepository.FindTagByName(ctx, blogDTO.Tag.Name)
		if err != nil {
			return blog, err
		}
		blogTag := entity.BlogTag{
			TagID: tag.ID,
			BlogID: blogCreate.ID,
		}
		_, err = bs.blogTagRepository.CreateBlogTag(ctx, blogTag)
		if err != nil {
			return blog, err
		}
	}
	return blogCreate, nil
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

func(bs *blogService) AssignTag(ctx context.Context, blogTagDTO dto.BlogTagCreateDto) (error) {
	blogUUID, _ := uuid.Parse(blogTagDTO.BlogID)
	tagUUID, _ := uuid.Parse(blogTagDTO.TagID)
	blogTag := entity.BlogTag{
		BlogID: blogUUID,
		TagID: tagUUID,
	}
	_, error := bs.blogTagRepository.CreateBlogTag(ctx, blogTag)
	return error
}