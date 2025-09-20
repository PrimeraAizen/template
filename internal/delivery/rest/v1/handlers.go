package v1

import (
	"github.com/gin-gonic/gin"

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

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	h.InitExampleRoutes(v1)
	h.InitHealthRoutes(v1)
}
