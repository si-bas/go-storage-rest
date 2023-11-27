package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/si-bas/go-storage-rest/domain/model"
	"github.com/si-bas/go-storage-rest/pkg/logger"
	"github.com/si-bas/go-storage-rest/pkg/logger/tag"
	"github.com/si-bas/go-storage-rest/shared/helper/pagination"
	"github.com/si-bas/go-storage-rest/shared/helper/response"
)

func (h *Handler) ClientRegister(c *gin.Context) {
	ctx := c.Request.Context()
	result := response.NewJSONResponse(ctx)

	var payload model.ClientCreatePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.Warn(ctx, "failed to bind JSON", tag.Err(err))
		c.JSON(result.APIStatusBadRequest().StatusCode, result.SetError(response.ErrBadRequest, err.Error()))
		return
	}

	isExists, err := h.clientService.IsExists(ctx, model.ClientExists{
		Name: payload.Name,
	})
	if err != nil {
		c.JSON(result.APIInternalServerError().StatusCode, result.SetError(response.ErrInternalServerError, err.Error()))
		return
	}

	if isExists {
		c.JSON(result.APIStatusConflict().StatusCode, result.SetError(response.ErrConflict, "client already exists"))
		return
	}

	client, err := h.clientService.Register(ctx, payload)
	if err != nil {
		c.JSON(result.APIInternalServerError().StatusCode, result.SetError(response.ErrInternalServerError, err.Error()))
		return
	}

	c.JSON(result.APIStatusCreated().StatusCode, result.SetData(client))
}

func (h *Handler) ClientList(c *gin.Context) {
	ctx := c.Request.Context()
	result := response.NewJSONResponse(ctx)

	var query model.ClientFilterParams
	if err := c.ShouldBindQuery(&query); err != nil {
		logger.Warn(ctx, "failed to bind query", tag.Err(err))
		c.JSON(result.APIStatusBadRequest().StatusCode, result.SetError(response.ErrBadRequest, err.Error()))
		return
	}

	var sortBys []pagination.ParamSort
	if len(query.Sort) > 0 {
		for k, v := range query.Sort {
			sortBys = append(sortBys, pagination.ParamSort{
				Column: k,
				Order:  v,
			})
		}
	}

	data, meta, err := h.clientService.GetListWithPagination(ctx, model.ClientFilter{
		Name: query.Name,
	}, pagination.Param{
		Limit: query.Limit,
		Page:  query.Page,
		Sort:  sortBys,
	})
	if err != nil {
		c.JSON(result.APIInternalServerError().StatusCode, result.SetError(response.ErrInternalServerError, err.Error()))
		return
	}

	c.JSON(result.APIStatusSuccess().StatusCode, result.SetData(data).SetMeta(meta))
}
