package service

import (
	"context"

	"github.com/si-bas/go-storage-rest/domain/model"
	"github.com/si-bas/go-storage-rest/domain/repository"
	"github.com/si-bas/go-storage-rest/pkg/logger"
	"github.com/si-bas/go-storage-rest/shared"
	"github.com/si-bas/go-storage-rest/shared/helper/pagination"
	"gorm.io/gorm"
)

type ClientService interface {
	GetOneByKey(ctx context.Context, key string) (*model.Client, error)
	GetListWithPagination(ctx context.Context, filter model.ClientFilter, pagination pagination.Param) ([]model.Client, *pagination.Param, error)

	IsExists(ctx context.Context, attr model.ClientExists) (bool, error)
	Register(ctx context.Context, payload model.ClientCreatePayload) (*model.Client, error)
}

type clientSrvImpl struct {
	clientRepo repository.ClientRepository
}

func NewClientService(clientRepo repository.ClientRepository) ClientService {
	return &clientSrvImpl{
		clientRepo: clientRepo,
	}
}

func (s *clientSrvImpl) GetOneByKey(ctx context.Context, key string) (*model.Client, error) {
	client, err := s.clientRepo.FindByKey(ctx, key)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		logger.Error(ctx, "failed to find client by key", err)
		return nil, err
	}

	return client, nil
}

func (s *clientSrvImpl) Register(ctx context.Context, payload model.ClientCreatePayload) (*model.Client, error) {

	var apiKey string

	for {
		key, err := shared.GenerateAPIKey()
		if err != nil {
			logger.Error(ctx, "failed to generate api key", err)
			return nil, err
		}

		keyCount, err := s.clientRepo.CountByKey(ctx, key)
		if err != nil {
			logger.Error(ctx, "failed to check if api key is exists", err)
			return nil, err
		}

		if *keyCount == 0 {
			apiKey = key
			break
		}
	}

	newClient := model.Client{
		Code: shared.GenerateSlug(payload.Name),
		Name: payload.Name,
		Key:  apiKey,
	}
	if err := s.clientRepo.Insert(ctx, &newClient); err != nil {
		logger.Error(ctx, "failed to insert client", err)
		return nil, err
	}

	return &newClient, nil
}

func (s *clientSrvImpl) IsExists(ctx context.Context, attr model.ClientExists) (bool, error) {
	if attr.Key != "" {
		count, err := s.clientRepo.CountByKey(ctx, attr.Key)
		if err != nil {
			logger.Error(ctx, "failed to count by key", err)
			return false, err
		}

		if *count > 0 {
			return true, nil
		}
	}

	if attr.Name != "" {
		count, err := s.clientRepo.CountByName(ctx, attr.Name)
		if err != nil {
			logger.Error(ctx, "failed to count by name", err)
			return false, err
		}

		if *count > 0 {
			return true, nil
		}
	}

	return false, nil
}

func (s *clientSrvImpl) GetListWithPagination(ctx context.Context, filter model.ClientFilter, pagination pagination.Param) ([]model.Client, *pagination.Param, error) {
	clients, meta, err := s.clientRepo.GetWithPagination(ctx, filter, pagination)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			logger.Error(ctx, "failed to get branch data with filter", err)
		}

		return nil, nil, err
	}

	return clients, meta, nil
}
