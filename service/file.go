package service

import (
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/si-bas/go-storage-rest/domain/model"
	"github.com/si-bas/go-storage-rest/domain/repository"
	"github.com/si-bas/go-storage-rest/pkg/logger"
	"github.com/si-bas/go-storage-rest/shared"
	"github.com/si-bas/go-storage-rest/shared/constant"
)

type FileService interface {
	GetUploadDir(ctx context.Context) (string, error)
	GenerateFilename(ctx context.Context, ext string) string

	Create(ctx context.Context, payload model.FileCreate) (*model.File, error)
}

type fileSrvImpl struct {
	fileRepo repository.FileRepository
}

func NewFileService(fileRepo repository.FileRepository) FileService {
	return &fileSrvImpl{
		fileRepo: fileRepo,
	}
}

func (s *fileSrvImpl) GenerateFilename(ctx context.Context, ext string) string {
	return uuid.New().String() + ext
}

func (s *fileSrvImpl) GetUploadDir(ctx context.Context) (string, error) {
	uploadDir := constant.UploadDir
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		logger.Error(ctx, "failed to create upload directory", err)
		return "", err
	}

	return uploadDir, nil
}

func (s *fileSrvImpl) Create(ctx context.Context, payload model.FileCreate) (*model.File, error) {
	clientID, err := shared.GetContextValueAsNumber(ctx, constant.XClientIDHeader)
	if err != nil {
		logger.Error(ctx, "failed to get client's ID from context", err)
		return nil, err
	}

	newFile := model.File{
		ClientID:     clientID,
		Code:         uuid.New().String(),
		OriginalName: payload.OriginalName,
		Name:         payload.Name,
		Extension:    payload.Extension,
		Size:         payload.Size,
		Path:         payload.Path,
	}
	if err := s.fileRepo.Insert(ctx, &newFile); err != nil {
		logger.Error(ctx, "failed to insert file", err)
		return nil, err
	}

	return &newFile, nil
}
