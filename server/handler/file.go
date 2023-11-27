package handler

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/si-bas/go-storage-rest/domain/model"
	"github.com/si-bas/go-storage-rest/pkg/logger"
	"github.com/si-bas/go-storage-rest/pkg/logger/tag"
	"github.com/si-bas/go-storage-rest/shared/helper/response"
)

func (h *Handler) FileUpload(c *gin.Context) {
	ctx := c.Request.Context()
	result := response.NewJSONResponse(ctx)

	uploadedFile, err := c.FormFile("file")
	if err != nil {
		logger.Warn(ctx, "failed to get uploaded file", tag.Err(err))
		c.JSON(result.APIStatusBadRequest().StatusCode, result.SetError(response.ErrBadRequest, err.Error()))
		return
	}

	uploadDir, err := h.fileService.GetUploadDir(ctx)
	if err != nil {
		c.JSON(result.APIInternalServerError().StatusCode, result.SetError(response.ErrInternalServerError, err.Error()))
		return
	}

	filename := uploadedFile.Filename
	fileExtension := filepath.Ext(filename)

	newFilename := h.fileService.GenerateFilename(ctx, fileExtension)
	if err := c.SaveUploadedFile(uploadedFile, filepath.Join(uploadDir, newFilename)); err != nil {
		logger.Error(ctx, "failed to move uploaded file to upload dir", err)
		c.JSON(result.APIInternalServerError().StatusCode, result.SetError(response.ErrInternalServerError, err.Error()))
		return
	}

	file, err := h.fileService.Create(ctx, model.FileCreate{
		OriginalName: filename,
		Name:         newFilename,
		Extension:    fileExtension,
		Size:         uploadedFile.Size,
		Path:         uploadDir,
	})
	if err != nil {
		c.JSON(result.APIInternalServerError().StatusCode, result.SetError(response.ErrInternalServerError, err.Error()))
		return
	}

	c.JSON(result.APIStatusCreated().StatusCode, result.SetData(file))
}
