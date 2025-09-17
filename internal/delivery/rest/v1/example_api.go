package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *Handler) InitExampleRoutes(router *gin.RouterGroup) {
	exampleRoutes := router.Group("/example")
	{
		exampleRoutes.GET("/", api.ExampleEndpoint)
	}
}

func (api *Handler) ExampleEndpoint(c *gin.Context) {
	err := api.services.ExampleService.ExampleMethod()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (api *Handler) InitHealthRoutes(router *gin.RouterGroup) {
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	router.GET("/readyz", func(c *gin.Context) {
		if err := api.services.HealthService.Ping(c.Request.Context()); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "not ready", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ready"})
	})
}
