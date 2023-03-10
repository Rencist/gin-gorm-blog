package repository

import (
	"context"
	"gin-gorm-blog/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlogTagRepository interface {
	CreateBlogTag(ctx context.Context, blogTag entity.BlogTag) (entity.BlogTag, error)
}

type blogTagConnection struct {
	connection *gorm.DB
}

func NewBlogTagRepository(db *gorm.DB) BlogTagRepository {
	return &blogTagConnection{
		connection: db,
	}
}

func(db *blogTagConnection) CreateBlogTag(ctx context.Context, blogTag entity.BlogTag) (entity.BlogTag, error) {
	blogTag.ID = uuid.New()
	btc := db.connection.Create(&blogTag)
	if btc.Error != nil {
		return entity.BlogTag{}, btc.Error
	}
	return blogTag, nil
}