package api

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up QuantumSynth's full API surface.
func RegisterRoutes(router *gin.Engine) {
	// Core system endpoints
	router.GET("/health", HealthCheckHandler)
	router.GET("/metrics", MetricsHandler)

	// Quantum/AI endpoints
	api := router.Group("/api/v1")
	{
		api.POST("/quantum/process", QuantumProcessHandler)
		api.GET("/models", ListModelsHandler)
		api.POST("/inference", InferenceHandler)
		// Add more ultra high-tech endpoints here!
	}
}
