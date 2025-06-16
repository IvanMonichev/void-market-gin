package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
	"time"
)

type Config struct {
	Env      string      `yaml:"env" env-default:"local"`
	Server   ServerCfg   `yaml:"server"`
	Postgres PostgresCfg `yaml:"postgres"`
	RabbitMQ RabbitMQCfg `yaml:"rabbitmq"`
}

type ServerCfg struct {
	Address string `yaml:"address" env-default:"0.0.0.0"`
	Port    string `yaml:"port" env-required:"true"`
}

type PostgresCfg struct {
	DSN     string        `yaml:"dsn" env-required:"true"`
	Timeout time.Duration `yaml:"timeout" env-default:"10s"`
}

type RabbitMQCfg struct {
	URL   string `yaml:"url" env-required:"true"`
	Queue string `yaml:"queue" env-default:"order_status_changed"`
}

func substitutePlaceholders(s string) string {
	for {
		start := strings.Index(s, "{")
		end := strings.Index(s, "}")
		if start == -1 || end == -1 || end < start {
			break
		}
		key := s[start+1 : end]
		val := os.Getenv(key)
		s = strings.Replace(s, "{"+key+"}", val, 1)
	}
	return s
}

func MustLoad() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found or failed to load it")
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set in environment")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	cfg.Server.Port = substitutePlaceholders(cfg.Server.Port)
	cfg.Postgres.DSN = substitutePlaceholders(cfg.Postgres.DSN)
	cfg.RabbitMQ.URL = substitutePlaceholders(cfg.RabbitMQ.URL)

	return &cfg
}
