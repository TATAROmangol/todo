package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HttpServer  `yaml:"http_server"`
}

type HttpServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080""`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"4s"`
}

var cfgPath = "./configs/config.yaml"

func MustLoad() *Config {
	file, err := os.ReadFile(cfgPath)
	if os.IsNotExist(err) {
		log.Fatalf("Config file not found: %s", cfgPath)
	}

	var cfg *Config
	if err = yaml.Unmarshal(file, &cfg); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}

	return cfg
}
