package synth

import (
	"math"
	"math/rand"
	"time"
)

// AdvancedQuantumNoise simulates quantum noise for stochastic processes.
func AdvancedQuantumNoise(seed int64) float64 {
	rand.Seed(seed)
	return rand.NormFloat64() * math.Sin(float64(seed)%math.Pi)
}

// QuantumRandomMatrix generates a matrix with quantum-like randomness.
func QuantumRandomMatrix(size int) [][]float64 {
	rand.Seed(time.Now().UnixNano())
	matrix := make([][]float64, size)
	for i := range matrix {
		matrix[i] = make([]float64, size)
		for j := range matrix[i] {
			matrix[i][j] = AdvancedQuantumNoise(time.Now().UnixNano() + int64(i*j))
		}
	}
	return matrix
}

// QuantumCollapse simulates a quantum collapse event on input data.
func QuantumCollapse(data []float64, threshold float64) []float64 {
	collapsed := make([]float64, len(data))
	for i, v := range data {
		if math.Abs(v) > threshold {
			collapsed[i] = 1
		} else {
			collapsed[i] = 0
		}
	}
	return collapsed
}
