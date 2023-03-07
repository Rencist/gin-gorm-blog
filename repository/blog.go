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