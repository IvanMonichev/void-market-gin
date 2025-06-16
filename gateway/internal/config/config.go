package config

import (
	"gateway/pkg/util"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Env      string      `yaml:"env"`
	Server   ServerCfg   `yaml:"server"`
	Services ServicesCfg `yaml:"services"`
}

type ServerCfg struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}

type ServicesCfg struct {
	User    string `yaml:"user"`
	Order   string `yaml:"order"`
	Payment string `yaml:"payment"`
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
	cfg.Services.User = util.SubstitutePlaceholders(cfg.Services.User)
	cfg.Services.Order = util.SubstitutePlaceholders(cfg.Services.Order)
	cfg.Services.Payment = util.SubstitutePlaceholders(cfg.Services.Payment)

	return &cfg
}
