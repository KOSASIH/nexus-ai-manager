package util

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	mrand "math/rand"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// GenerateSecureID generates a cryptographically secure random hex string of n bytes.
func GenerateSecureID(n int) (string, error) {
	if n <= 0 {
		return "", errors.New("invalid length for secure ID")
	}
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("failed to generate secure ID: %w", err)
	}
	return hex.EncodeToString(b Retry retries the given function up to maxAttempts with exponential backoff.
func Retry(maxAttempts int, sleep time.Duration, fn func() error) error {
	var err error
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		err = fn()
		if err == nil {
			return nil
		}
		logrus.WithFields(logrus.Fields{
			"attempt": attempt,
			"error":   err,
		}).Warn("Retryable function failed")
		time.Sleep(time.Duration(math.Pow(2, float64(attempt-1))) * sleep)
	}
	return fmt.Errorf("all %d attempts failed: %w", maxAttempts, err)
}

// GetEnv fetches an environment variable or returns the fallback value.
func GetEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

// TimeTrack logs the duration of a function (use with defer).
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	logrus.WithFields(logrus.Fields{
		"function": name,
		"elapsed":  elapsed,
	}).Info("Execution time tracked")
}

// SafeGo runs a goroutine with panic recovery and error logging.
func SafeGo(name string, fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				logrus.WithFields(logrus.Fields{
					"goroutine": name,
					"panic":     r,
					"stack":     string(StackTrace()),
				}).Error("Panic recovered in goroutine")
			}
		}()
		fn()
	}()
}

// StackTrace returns the current goroutine stack trace.
func StackTrace() []byte {
	buf := make([]byte,16)
	runtime.Stack(buf, false)
	return buf
}

// SanitizeString removes dangerous or unwanted characters from input.
func SanitizeString(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\x00", "")
	return s
}

// RandomIntInRange returns a random int in [min, max].
func RandomIntInRange(min, max int) int {
	if min >= max {
		return min
	}
	mrand.Seed(time.Now().UnixNano())
	return mrand.Intn(max-min+1) + min
}
