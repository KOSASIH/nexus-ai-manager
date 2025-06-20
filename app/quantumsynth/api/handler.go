package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"quantumsynth/internal/synth"
	"quantumsynth/model"
	"github.com/sirupsen/logrus"
)

// HealthCheckHandler provides a basic health check endpoint.
func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":     "ok",
		"service":    "QuantumSynth",
		"version":    "1.0.0",
		"timestamp":  time.Now().UTC(),
		"serverTime": time.Now().Format(time.RFC3339),
	})
}

// QuantumProcessHandler runs quantum-inspired data processing.
func QuantumProcessHandler(c *gin.Context) {
	var req model.ProcessRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.WithError(err).Warn("Invalid process request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "detail": err.Error()})
		return
	}

	result, err := synth.ProcessQuantumData(req)
	if err != nil {
		logrus.WithError(err).Error("Quantum processing failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Processing failed", "detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"result": result,
	})
}

// ListModelsHandler returns available quantum/AI models.
func ListModelsHandler(c *gin.Context) {
	models := synth.AvailableModels()
	c.JSON(http.StatusOK, gin.H{
		"models": models,
	})
}

// InferenceHandler runs an inference using a selected model.
func InferenceHandler(c *gin.Context) {
	var req model.InferenceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.WithError(err).Warn("Invalid inference request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "detail": err.Error()})
		return
	}
	output, err := synth.RunInference(req)
	if err != nil {
		logrus.WithError(err).Error("Inference failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Inference failed", "detail": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"output": output})
}

// MetricsHandler returns server metrics (stub for full Prometheus integration).
func MetricsHandler(c *gin.Context) {
	// For full production, integrate with Prometheus or OpenTelemetry exporters.
	c.JSON(http.StatusOK, gin.H{
		"metrics": map[string]interface{}{
			"uptime_seconds":   synth.Uptime(),
			"requests_total":   synth.RequestCount(),
			"active_sessions":  synth.ActiveSessions(),
			"quantum_jobs_run": synth.QuantumJobCount(),
		},
	})
}
