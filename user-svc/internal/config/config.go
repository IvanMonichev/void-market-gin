package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env    string    `yaml:"env" env-default:"local"`
	Server ServerCfg `yaml:"server"`
	Mongo  MongoCfg  `yaml:"mongo"`
}

type ServerCfg struct {
	Address string `yaml:"address" env-default:"0.0.0.0"`
	Port    string `yaml:"port" env-required:"true"`
}

type MongoCfg struct {
	URI      string        `yaml:"uri" env-required:"true"`
	Database string        `yaml:"database" env-default:"void_market_user"`
	Timeout  time.Duration `yaml:"timeout" env-default:"10s"`
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
		log.Fatal("Environment variable CONFIG_PATH is not set")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	// Подстановка переменных окружения
	cfg.Server.Port = substitutePlaceholders(cfg.Server.Port)
	cfg.Mongo.URI = substitutePlaceholders(cfg.Mongo.URI)

	return &cfg
}
