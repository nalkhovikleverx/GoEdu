package config

import (
	"net"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/stackus/dotenv"
)

type HTTP struct {
	Host string `default:"0.0.0.0"`
	Port string `default:"8080"`
}

func (h HTTP) Address() string {
	return net.JoinHostPort(h.Host, h.Port)
}

// Config is the type that holds the configuration for the service.
type Config struct {
	ServiceName     string `envconfig:"SERVICE_NAME" default:"noname"`
	Environment     string `envconfig:"ENVIRONMENT" default:"dev"`
	LogLevel        string `envconfig:"LOG_LEVEL" default:"DEBUG"`
	HTTP            HTTP
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
}

func InitConfig() (*Config, error) {
	cfg := Config{}
	err := dotenv.Load(dotenv.EnvironmentFiles(os.Getenv("ENVIRONMENT")))
	if err != nil {
		return nil, err
	}
	return &cfg, envconfig.Process("", &cfg)
}
