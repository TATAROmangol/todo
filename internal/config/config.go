package config

import (
	"fmt"
	"log"
	"os"
	v1 "todo/internal/servers/http/v1"
	"todo/pkg/sqlite"
)

type Config struct {
	HttpConfig v1.Config
	RepoConfig sqlite.Config
}

func MustLoad() Config {
	httpPort, exist := os.LookupEnv("HTTP_PORT")
	if !exist {
		log.Fatal("no found env HTTP_PORT")
	}
	httpAddress := fmt.Sprintf("localhost:%v", httpPort)

	repoPath, exist := os.LookupEnv("STORAGE_PATH")
	if !exist {
		log.Fatal("no found env HTTP_PORT")
	}

	return Config{
		HttpConfig: v1.Config{
			Address: httpAddress,
		},
		RepoConfig: sqlite.Config{
			Path: repoPath,
		},
	}
}
