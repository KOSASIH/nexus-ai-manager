package config

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Server   ServerConfig   `mapstructure:"server"`
	Log      LogConfig      `mapstructure:"log"`
	Quantum  QuantumConfig  `mapstructure:"quantum"`
	Security SecurityConfig `mapstructure:"security"`
}

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
}

type QuantumConfig struct {
	DefaultMode string `mapstructure:"default_mode"`
	MaxJobs     int    `mapstructure:"max_jobs"`
}

type SecurityConfig struct {
	EnableAuth      bool   `mapstructure:"enable_auth"`
	JWTSecret       string `mapstructure:"jwt_secret"`
	AllowedOrigins  string `mapstructure:"allowed_origins"`
}

var (
	once      sync.Once
	appConfig *AppConfig
)

// LoadConfig loads app configuration (singleton).
func LoadConfig(configPath string) (*AppConfig, error) {
	var err error
	once.Do(func() {
		viper.SetConfigType("yaml")
		if configPath != "" {
			viper.SetConfigFile(configPath)
		} else {
			viper.SetConfigName("config")
			viper.AddConfigPath(".")
			viper.AddConfigPath("./config")
		}

		// Env var override support
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.AutomaticEnv()

		if err = viper.ReadInConfig(); err != nil {
			logrus.Warnf("Config file not found, using ENV or defaults: %v", err)
		}
		appConfig = &AppConfig{}
		if err = viper.Unmarshal(&appConfig); err != nil {
			err = fmt.Errorf("unable to decode config into struct: %w", err)
			appConfig = nil
			return
		}
		applyDefaults(appConfig)
	})
	return appConfig, err
}

// applyDefaults ensures sane defaults for any missing config values.
func applyDefaults(cfg *AppConfig) {
	if cfg.Server.Host == "" {
		cfg.Server.Host = "0.0.0.0"
	}
	if cfg.Server.Port == 0 {
		cfg.Server.Port = 8080
	}
	if cfg.Log.Level == "" {
		cfg.Log.Level = "info"
	}
	if cfg.Quantum.DefaultMode == "" {
		cfg.Quantum.DefaultMode = "superposition"
	}
	if cfg.Quantum.MaxJobs == 0 {
		cfg.Quantum.MaxJobs = 64
	}
	if cfg.Security.JWTSecret == "" {
		cfg.Security.JWTSecret = os.Getenv("QUANTUMSYNTH_JWT_SECRET")
	}
}

// GetConfig returns the loaded config, or panics if uninitialized.
func GetConfig() *AppConfig {
	if appConfig == nil {
		panic("Config not loaded. Call LoadConfig first!")
	}
	return appConfig
}
