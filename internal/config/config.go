package config

import (
	"log"
	"log/slog"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env       string     `yaml:"env" env-default:"local"`
	MongoPath string     `yaml:"mongo_path" env-default:"mongodb://localhost:27017" env-required:"true"`
	MongoDB   string     `yaml:"mongo_db" env-default:"shelter" env-required:"true"`
	HTTP      HTTPConfig `yaml:"http_server" env-required:"true"`
}

type HTTPConfig struct {
	Port         string `yaml:"port" env-default:"8080" env-required:"true"`
	Timeout      int    `yaml:"timeout" env-default:"5" env-required:"true"`
	IddleTimeout int    `yaml:"iddle_timeout" env-default:"10" env-required:"true"`
}

func NewConfig() Config {
	cfg := Config{}
	err := cleanenv.ReadConfig("internal/config/config.yaml", &cfg)
	if err != nil {
		log.Fatalf("Configuration cannot be read: %v", err)
	}
	return cfg
}

func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "prod":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}
