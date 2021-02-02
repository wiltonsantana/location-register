package config

import (
	"strings"

	"github.com/wiltonsantana/location-register/pkg/logging"

	"github.com/spf13/viper"
)

// Server represents the server configuration properties
type Server struct {
	Port int
}

// Logger represents the logger configuration properties
type Logger struct {
	Level  string
	Syslog bool
}

// Config represents the service configuration
type Config struct {
	Server
	Logger
}

func readFile(name string) {
	logger := logging.NewLogrus("error", false).Get("Config")
	viper.SetConfigName(name)
	if err := viper.ReadInConfig(); err != nil {
		logger.Fatalf("error reading config file, %s", err)
	}
}

// Load returns the service configuration
func Load() Config {
	var configuration Config
	logger := logging.NewLogrus("error", false).Get("Config")
	viper.AddConfigPath("internal/config")
	viper.SetConfigType("yaml")

	readFile("default")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	if err := viper.Unmarshal(&configuration); err != nil {
		logger.Fatalf("error unmarshalling configuration, %s", err)
	}

	return configuration
}
