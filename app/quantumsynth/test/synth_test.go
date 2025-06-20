package test

import (
	"math"
	"testing"
	"time"

	"quantumsynth/internal/synth"
	"quantumsynth/model"
)

// TestProcessQuantumData covers core quantum-inspired processing.
func TestProcessQuantumData(t *testing.T) {
	cases := []struct {
		name    string
		input   string
		mode    string
		wantErr bool
	}{
		{"Superposition", "testdata", "superposition", false},
		{"Entanglement", "qbit", "entanglement", false},
		{"QuantumWalk", "walker", "quantum-walk", false},
		{"ClassicFallback", "legacy", "classic", false},
		{"EmptyInput", "", "superposition", true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := model.ProcessRequest{Input: c.input, Mode: c.mode}
			res, err := synth.ProcessQuantumData(req)
			if (err != nil) != c.wantErr {
				t.Fatalf("expected error: %v, got: %v (err: %v)", c.wantErr, err == nil, err)
			}
			if !c.wantErr && res.Output == "" {
				t.Error("expected non-empty output")
			}
			if !c.wantErr && res.Timestamp.After(time.Now().Add(1*time.Minute)) {
				t.Error("unexpected future timestamp")
			}
		})
	}
}

// TestAvailableModels ensures model listing is correct.
func TestAvailableModels(t *testing.T) {
	models := synth.AvailableModels()
	expected := []string{"superposition", "entanglement", "quantum-walk", "deep-neuro-synth"}
	if len(models) != len(expected) {
		t.Fatalf("expected %v models, got %v", len(expected), len(models))
	}
	for _, m := range expected {
		found := false
		for _, got := range models {
			if m == got {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("model %q not found in available models", m)
		}
	}
}

// TestRunInference checks inference logic.
func TestRunInference(t *testing.T) {
	req := model.InferenceRequest{Model: "superposition", Data: "test"}
	output, err := synth.RunInference(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if output == "" {
		t.Error("expected non-empty output")
	}
	_, err = synth.RunInference(model.InferenceRequest{})
	if err == nil {
		t.Error("expected error for empty input")
	}
}

// TestQuantumRandomMatrix verifies matrix randomness and structure.
func TestQuantumRandomMatrix(t *testing.T) {
	matrix := synth.QuantumRandomMatrix(4)
	if len(matrix) != 4 || len(matrix[0]) != 4 {
		t.Fatalf("expected 4x4 matrix, got %dx%d", len(matrix), len(matrix[0]))
	}
	// Check for non-zero variance
	sum := 0.0
	for i := range matrix {
		for j := range matrix[i] {
			sum += matrix[i][j]
		}
	}
	if math.Abs(sum) < 0.0001 {
		t.Error("matrix appears non-random or all zeros")
	}
}

// TestQuantumCollapse checks collapse logic.
func TestQuantumCollapse(t *testing.T) {
	data := []float64{0.1, -2.5, 0.5, 3.0}
	result := synth.QuantumCollapse(data, 1.0)
	want := []float64{0,	if len(result) != len(want) {
		t.Fatalf("expected len %d, got %d", len(want), len(result))
	}
	for i := range want {
		if result[i] != want[i] {
			t.Errorf("collapse mismatch at %d: want %v, got %v", i, want[i], result[i])
		}
	}
}

// TestAdvancedQuantumNoise checks noise generation.
func TestAdvancedQuantumNoise(t *testing.T) {
	val := synth.AdvancedQuantumNoise(42)
	if math.IsNaN(val) {
		t.Error("expected a real value, got NaN")
	}
}

// TestSafeConcurrency ensures helpers can be used concurrently.
func TestSafeConcurrency(t *testing.T) {
	done := make(chan bool)
	for i := 0; i < 8; i++ {
		go func(i int) {
			_ = synth.AdvancedQuantumNoise(int64(i))
			done <- true
		}(i)
	}
	for i := 0; i < 8; i++ {
		<-done
	}
}
