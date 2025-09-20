package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PrimeraAizen/template/config"
	v1 "github.com/PrimeraAizen/template/internal/delivery/rest/v1"
	"github.com/PrimeraAizen/template/internal/service"
	"github.com/PrimeraAizen/template/pkg/logger"
)

type Handler struct {
	services *service.Service
	logger   *logger.Logger
}

func NewHandler(services *service.Service, appLogger *logger.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   appLogger,
	}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	router := gin.New()

	// Add custom middleware
	router.Use(
		logger.RequestIDMiddleware(),
		logger.LoggingMiddleware(h.logger),
		logger.RecoveryMiddleware(h.logger),
		logger.ContextMiddleware(h.logger),
	)

	// Health check endpoint
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services, h.logger)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
