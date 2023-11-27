package middleware

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/si-bas/go-storage-rest/config"
	"github.com/si-bas/go-storage-rest/service"
	"github.com/si-bas/go-storage-rest/shared/constant"
)

func AdminApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader(constant.XApiKeyHeader)
		if apiKey == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "API Key is required"})
			c.Abort()
			return
		}

		if apiKey != config.Config.Api.Key {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "API Key is incorrect"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func ClientApiKey(clientSrv service.ClientService) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader(constant.XApiKeyHeader)
		if apiKey == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "API key is required"})
			c.Abort()
			return
		}

		client, err := clientSrv.GetOneByKey(c.Request.Context(), apiKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Auth error"})
			c.Abort()
			return
		}

		if client == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid API key"})
			c.Abort()
			return
		}

		clientIDstr := strconv.FormatUint(client.ID, 10)
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), constant.XClientIDHeader, clientIDstr))

		c.Next()
	}
}
