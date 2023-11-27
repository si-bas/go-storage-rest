package repository

import (
	"context"

	"github.com/si-bas/go-storage-rest/domain/model"
	"gorm.io/gorm"
)

type FileRepository interface {
	Insert(ctx context.Context, file *model.File) error
}

type fileRepoImpl struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) FileRepository {
	return &fileRepoImpl{
		db: db,
	}
}

func (r *fileRepoImpl) Insert(ctx context.Context, file *model.File) error {
	return r.db.Model(&model.File{}).Create(file).Error
}
