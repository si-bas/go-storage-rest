package repository

import (
	"context"

	"github.com/si-bas/go-storage-rest/domain/model"
	"github.com/si-bas/go-storage-rest/shared/helper/pagination"
	"gorm.io/gorm"
)

type ClientRepository interface {
	Insert(ctx context.Context, client *model.Client) error
	FindByKey(ctx context.Context, key string) (*model.Client, error)

	CountByName(ctx context.Context, name string) (*int64, error)
	CountByKey(ctx context.Context, key string) (*int64, error)

	FilteredDb(filter model.ClientFilter) *gorm.DB
	GetWithPagination(ctx context.Context, filter model.ClientFilter, param pagination.Param) ([]model.Client, *pagination.Param, error)
}

type clientRepoImpl struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) ClientRepository {
	return &clientRepoImpl{
		db: db,
	}
}

func (r *clientRepoImpl) Insert(ctx context.Context, client *model.Client) error {
	return r.db.Model(&model.Client{}).Create(client).Error
}

func (r *clientRepoImpl) FindByKey(ctx context.Context, key string) (*model.Client, error) {
	var client model.Client

	if err := r.db.Model(&model.Client{}).Where("`key` = ?", key).First(&client).Error; err != nil {
		return nil, err
	}

	return &client, nil
}

func (r *clientRepoImpl) CountByName(ctx context.Context, name string) (*int64, error) {
	var count int64

	if err := r.db.Model(&model.Client{}).Where("`name` = ?", name).Count(&count).Error; err != nil {
		return nil, err
	}

	return &count, nil
}

func (r *clientRepoImpl) CountByKey(ctx context.Context, key string) (*int64, error) {
	var count int64

	if err := r.db.Model(&model.Client{}).Where("`key` = ?", key).Count(&count).Error; err != nil {
		return nil, err
	}

	return &count, nil
}

func (r *clientRepoImpl) FilteredDb(filter model.ClientFilter) *gorm.DB {
	chain := r.db.Model(&model.Client{})

	if filter.Name != "" {
		chain.Where("`name` LIKE ?", "%"+filter.Name+"%")
	}

	return chain
}

func (r *clientRepoImpl) GetWithPagination(ctx context.Context, filter model.ClientFilter, param pagination.Param) ([]model.Client, *pagination.Param, error) {
	var clients []model.Client

	filteredDb := r.FilteredDb(filter)
	if err := filteredDb.Scopes(pagination.Paginate(model.Client{}, &param, filteredDb)).Find(&clients).Error; err != nil {
		return nil, nil, err
	}

	return clients, &param, nil
}
