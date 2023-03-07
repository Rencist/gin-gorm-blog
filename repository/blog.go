package repository

import (
	"context"
	"gin-gorm-blog/common"
	"gin-gorm-blog/dto"
	"gin-gorm-blog/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlogRepository interface {
	GetTotalData(ctx context.Context) (int64, error)
	CreateBlog(ctx context.Context, blog entity.Blog) (entity.Blog, error)
	GetAllBlog(ctx context.Context) ([]entity.Blog, error)
	FindBlogByUserID(ctx context.Context, userID string) ([]entity.Blog, error)
	FindBlogByID(ctx context.Context, blogID string) (entity.Blog, error)
	UpdateBlog(ctx context.Context, blog entity.Blog) (error)
	LikeBlogByID(ctx context.Context, blogID string) (error)
	CheckBlogCommentByID(ctx context.Context, blogID string) (entity.Blog, error)
	GetAllBlogPagination(ctx context.Context, pagination entity.Pagination) (dto.PaginationResponse, error)
}

type blogConnection struct {
	connection *gorm.DB
}

func NewBlogRepository(db *gorm.DB) BlogRepository {
	return &blogConnection{
		connection: db,
	}
}

func (db *blogConnection) GetTotalData(ctx context.Context) (int64, error) {
	var totalData int64
	bc := db.connection.Model(&entity.Blog{}).Count(&totalData)
	if bc.Error != nil {
		return 0, bc.Error
	}
	return totalData, nil
}

func(db *blogConnection) CreateBlog(ctx context.Context, blog entity.Blog) (entity.Blog, error) {
	blog.ID = uuid.New()
	blog.LikeCount = 0
	blog.WatchCount = 0
	bc := db.connection.Create(&blog)
	if bc.Error != nil {
		return entity.Blog{}, bc.Error
	}
	return blog, nil
}

func(db *blogConnection) GetAllBlog(ctx context.Context) ([]entity.Blog, error) {
	var listBlog []entity.Blog
	bc := db.connection.Find(&listBlog)
	if bc.Error != nil {
		return nil, bc.Error
	}
	return listBlog, nil
}

func(db *blogConnection) FindBlogByUserID(ctx context.Context, userID string) ([]entity.Blog, error) {
	var listBlog []entity.Blog
	bc := db.connection.Where("user_id = ?", userID).Find(&listBlog)
	if bc.Error != nil {
		return []entity.Blog{}, nil
	}
	return listBlog, nil
}

func(db *blogConnection) FindBlogByID(ctx context.Context, blogID string) (entity.Blog, error) {
	var blog entity.Blog
	bc := db.connection.Preload("Comments").Where("id = ?", blogID).Find(&blog)
	if bc.Error != nil {
		return entity.Blog{}, bc.Error
	}
	blog.WatchCount = blog.WatchCount + 1
	db.UpdateBlog(ctx, blog)
	return blog, nil
}

func(db *blogConnection) CheckBlogCommentByID(ctx context.Context, blogID string) (entity.Blog, error) {
	var blog entity.Blog
	bc := db.connection.Preload("Comments").Where("id = ?", blogID).Find(&blog)
	if bc.Error != nil {
		return entity.Blog{}, bc.Error
	}
	return blog, nil
}

func(db *blogConnection) UpdateBlog(ctx context.Context, blog entity.Blog) (error) {
	bc := db.connection.Updates(&blog)
	if bc.Error != nil {
		return bc.Error
	}
	return nil
}

func(db *blogConnection) LikeBlogByID(ctx context.Context, blogID string) (error) {
	var blog entity.Blog
	bc := db.connection.Where("id = ?", blogID).Find(&blog)
	if bc.Error != nil {
		return bc.Error
	}
	blog.LikeCount = blog.LikeCount + 1
	db.UpdateBlog(ctx, blog)
	return nil
}

func (db *blogConnection) GetAllBlogPagination(ctx context.Context, pagination entity.Pagination) (dto.PaginationResponse, error) {
	var paginationResponse dto.PaginationResponse
	var blogList []*entity.Blog

	totalData, _ := db.GetTotalData(ctx)

	db.connection.Debug().Scopes(common.PaginationOffset(&pagination, totalData)).Find(&blogList)
	pagination.DataPerPage = blogList
	paginationResponse.DataPerPage = blogList
	paginationResponse.Meta.MaxPage = pagination.MaxPage
	paginationResponse.Meta.Page = pagination.Page
	paginationResponse.Meta.TotalData = pagination.TotalData
	return paginationResponse, nil
}