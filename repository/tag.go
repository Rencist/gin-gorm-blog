package repository

import (
	"context"
	"gin-gorm-blog/entity"

	"gorm.io/gorm"
)

type TagRepository interface {
	FindTagByName(ctx context.Context, name string) (entity.Tag, error)
}

type tagConnection struct {
	connection *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return  &tagConnection{
		connection: db,
	}
}

func(db *tagConnection) FindTagByName(ctx context.Context, name string) (entity.Tag, error) {
	var tag entity.Tag
	tc := db.connection.Where("name = ?", name).Find(&tag)
	if tc.Error != nil {
		return entity.Tag{}, tc.Error
	}
	return tag, nil
}