package config

import (
	"log"
	"os"
	"time"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HttpServer `yaml:"http_server"`
}

type HttpServer struct {
	Address string `yaml:"address" env-default:"localhost:8080""`
	Timeout time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"4s"`
}

var cfgPath = "./config/config.yaml"

func MustLoad() *Config{
	if _, err := os.Stat(cfgPath); os.IsNotExist(err){
		log.Fatalf("Config file not found: %s", cfgPath)
	}
	
	var cfg *Config
	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return cfg
}