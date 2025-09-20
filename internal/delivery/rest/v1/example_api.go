package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PrimeraAizen/template/pkg/logger"
)

func (api *Handler) InitExampleRoutes(router *gin.RouterGroup) {
	exampleRoutes := router.Group("/example")
	{
		exampleRoutes.GET("/", api.ExampleEndpoint)
	}
}

func (api *Handler) ExampleEndpoint(c *gin.Context) {
	// Get logger from context
	appLogger := logger.GetLoggerFromContext(c.Request.Context())

	appLogger.WithComponent("api").WithOperation("example_endpoint").Info("Processing example request")

	err := api.services.ExampleService.ExampleMethod()
	if err != nil {
		appLogger.WithComponent("api").WithOperation("example_endpoint").WithError(err).Error("Example method failed")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":      err.Error(),
			"request_id": c.GetString("request_id"),
		})
		return
	}

	appLogger.WithComponent("api").WithOperation("example_endpoint").Info("Example request completed successfully")
	c.JSON(http.StatusOK, gin.H{
		"status":     "ok",
		"request_id": c.GetString("request_id"),
	})
}

func (api *Handler) InitHealthRoutes(router *gin.RouterGroup) {
	router.GET("/healthz", func(c *gin.Context) {
		appLogger := logger.GetLoggerFromContext(c.Request.Context())
		appLogger.WithComponent("health").WithOperation("healthz").Debug("Health check requested")
		c.JSON(http.StatusOK, gin.H{
			"status":     "ok",
			"request_id": c.GetString("request_id"),
		})
	})

	router.GET("/readyz", func(c *gin.Context) {
		appLogger := logger.GetLoggerFromContext(c.Request.Context())
		appLogger.WithComponent("health").WithOperation("readyz").Debug("Readiness check requested")

		if err := api.services.HealthService.Ping(c.Request.Context()); err != nil {
			appLogger.WithComponent("health").WithOperation("readyz").WithError(err).Error("Readiness check failed")
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status":     "not ready",
				"error":      err.Error(),
				"request_id": c.GetString("request_id"),
			})
			return
		}

		appLogger.WithComponent("health").WithOperation("readyz").Debug("Readiness check passed")
		c.JSON(http.StatusOK, gin.H{
			"status":     "ready",
			"request_id": c.GetString("request_id"),
		})
	})
}
