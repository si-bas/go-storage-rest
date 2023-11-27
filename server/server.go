package server

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/si-bas/go-storage-rest/config"
	"github.com/si-bas/go-storage-rest/domain/repository"
	"github.com/si-bas/go-storage-rest/pkg/gorm"
	"github.com/si-bas/go-storage-rest/pkg/logger"
	"github.com/si-bas/go-storage-rest/server/handler"
	"github.com/si-bas/go-storage-rest/server/middleware"
	"github.com/si-bas/go-storage-rest/service"
	"github.com/si-bas/go-storage-rest/shared/constant"
)

type HTTPServer struct {
}

// New to instantiate HTTPServer
func New() *HTTPServer {
	return &HTTPServer{}
}

func (s *HTTPServer) Start() {
	h := initHandler()

	if config.Config.App.Env == constant.EnvProduction {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	if config.Config.App.Env != constant.EnvProduction {
		router.Use(middleware.CORS())
	}

	router.Use(middleware.InjectContext())
	router.GET("/health", h.HealthCheck)

	router.Static("/f", constant.UploadDir)

	groupV1 := router.Group("/v1")

	admV1 := groupV1.Group("/adm")
	admV1.Use(middleware.AdminApiKey())

	groupV1.Use(middleware.ClientApiKey(h.ClientService))
	groupV1.POST("/file/upload", h.FileUpload)

	admV1.POST("/client", h.ClientRegister)
	admV1.GET("/client", h.ClientList)

	err := router.Run(fmt.Sprintf(":%d", config.Config.App.Port))
	if err != nil {
		logger.Error(context.Background(), "failed to run router", err)
	}
}

func initHandler() *handler.Handler {
	var err error
	config.TimeLocation, err = time.LoadLocation(config.Config.App.Timezone)
	if err != nil {
		panic("error set timezone, err=" + err.Error())
	}

	logger.InitLogger()

	// Init DB
	db := gorm.ConnectDB()

	// Init repositories
	clientRepo := repository.NewClientRepository(db)
	fileRepo := repository.NewFileRepository(db)

	// Init services
	clientSrv := service.NewClientService(clientRepo)
	fileSrv := service.NewFileService(fileRepo)

	return handler.New(clientSrv, fileSrv)
}
