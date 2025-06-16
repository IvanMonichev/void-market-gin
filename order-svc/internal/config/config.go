package config

import (
	"github.com/IvanMonichev/void-market-gin/order-svc/pkg/util"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
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

	cfg.Server.Port = util.SubstitutePlaceholders(cfg.Server.Port)
	cfg.Postgres.DSN = util.SubstitutePlaceholders(cfg.Postgres.DSN)
	cfg.RabbitMQ.URL = util.SubstitutePlaceholders(cfg.RabbitMQ.URL)

	return &cfg
}
