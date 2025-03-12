package config

import (
	v1 "todo/internal/servers/http/v1"
	"todo/pkg/postgres"
)

type Config struct {
	HttpConfig v1.Config
	RepoConfig postgres.Config
}

func MustLoad() Config {
	httpConfig := v1.MustLoadConfig()
	repoConfig := postgres.MustLoadConfig()

	return Config{
		HttpConfig: httpConfig,
		RepoConfig: repoConfig,
	}
}
