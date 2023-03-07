package repository

import (
	"context"
	"gin-gorm-blog/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlogRepository interface {
	CreateBlog(ctx context.Context, blog entity.Blog) (entity.Blog, error)
	GetAllBlog(ctx context.Context) ([]entity.Blog, error)
	FindBlogByUserID(ctx context.Context, userID string) ([]entity.Blog, error)
	FindBlogByID(ctx context.Context, blogID string) (entity.Blog, error)
	UpdateBlog(ctx context.Context, blog entity.Blog) (error)
	LikeBlogByID(ctx context.Context, blogID string) (error)
	CheckBlogCommentByID(ctx context.Context, blogID string) (entity.Blog, error)
}

type blogConnection struct {
	connection *gorm.DB
}

func NewBlogRepository(db *gorm.DB) BlogRepository {
	return &blogConnection{
		connection: db,
	}
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