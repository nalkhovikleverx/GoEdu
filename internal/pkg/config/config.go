package config

import (
	"log/slog"
	"net"
	"os"
	"runtime/debug"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/stackus/dotenv"
)

type HTTP struct {
	Host              string        `default:"0.0.0.0"`
	Port              string        `default:"8080"`
	ReadHeaderTimeout time.Duration `envconfig:"HTTP_READ_HEADER_TIMEOUT" default:"1s"`
}

func (h HTTP) Address() string {
	return net.JoinHostPort(h.Host, h.Port)
}

func (h HTTP) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("host", h.Host),
		slog.String("port", h.Port),
		slog.String("readHeaderTimeout", h.ReadHeaderTimeout.String()),
	)
}

type VCSInfo struct {
	Revision string
	Tag      string
	Time     string
}

// Config is the type that holds the configuration for the service.
type Config struct {
	VCS             VCSInfo
	ServiceName     string `envconfig:"SERVICE_NAME" default:"noname"`
	Environment     string `envconfig:"ENVIRONMENT" default:"dev"`
	LogLevel        string `envconfig:"LOG_LEVEL" default:"DEBUG"`
	HTTP            HTTP
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
}

const (
	vcsRevisionKey = "vcs.revision"
	vcsTagKey      = "vcs.tag"
	vcsTimeKey     = "vcs.time"
	gitShortOffset = 8
)

func InitConfig() (*Config, error) {
	cfg := Config{}
	info, _ := debug.ReadBuildInfo()
	for _, s := range info.Settings {
		switch s.Key {
		case vcsRevisionKey:
			cfg.VCS.Revision = s.Value[:gitShortOffset]
		case vcsTimeKey:
			cfg.VCS.Time = s.Value
		case vcsTagKey:
			cfg.VCS.Tag = s.Value
		}
	}
	err := dotenv.Load(dotenv.EnvironmentFiles(os.Getenv("ENVIRONMENT")))
	if err != nil {
		return nil, err
	}
	return &cfg, envconfig.Process("", &cfg)
}

func (c Config) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("serviceName", c.ServiceName),
		slog.String("env", c.Environment),
		slog.String("logLevel", c.LogLevel),
		slog.Any("http", c.HTTP),
		slog.String("shutdownTimeout", c.ShutdownTimeout.String()),
	)
}
