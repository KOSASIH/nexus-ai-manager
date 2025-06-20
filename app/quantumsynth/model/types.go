package model

import (
	"time"
)

// ProcessRequest defines the input for quantum-inspired data processing.
type ProcessRequest struct {
	Input string `json:"input" binding:"required"`
	Mode  string `json:"mode" binding:"required,oneof=superposition entanglement quantum-walk deep-neuro-synth"`
}

// ProcessResult defines the output/result of a quantum processing operation.
type ProcessResult struct {
	Output    string    `json:"output"`
	Timestamp time.Time `json:"timestamp"`
}

// InferenceRequest defines the input for running an AI/quantum model inference.
type InferenceRequest struct {
	Model string `json:"model" binding:"required"`
	Data  string `json:"data" binding:"required"`
}

// InferenceResult defines the output/result of an inference operation.
type InferenceResult struct {
	Output    string    `json:"output"`
	Model     string    `json:"model"`
	Timestamp time.Time `json:"timestamp"`
}

// ErrorResponse standardizes API error responses.
type ErrorResponse struct {
	Error  string `json:"error"`
	Detail string `json:"detail,omitempty"`
}

// QuantumMatrixRequest – request for quantum matrix generation.
type QuantumMatrixRequest struct {
	Size int `json:"size" binding:"required,min=2,max=128"`
}

// QuantumMatrixResult – result for a quantum matrix operation.
type QuantumMatrixResult struct {
	Matrix    [][]float64 `json:"matrix"`
	Timestamp time.Time   `json:"timestamp"`
}

// CollapseRequest – request for quantum collapse operation.
type CollapseRequest struct {
	Data      []float64 `json:"data" binding:"required"`
	Threshold float64   `json:"threshold" binding:"required"`
}

// CollapseResult – result of quantum collapse.
type CollapseResult struct {
	Collapsed []float64 `json:"collapsed"`
	Timestamp time.Time `json:"timestamp"`
}
