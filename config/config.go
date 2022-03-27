package config

import (
	"fmt"

	"enigmacamp.com/song/manager"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Config struct {
	InfraManager   manager.Infra
	RepoManager    manager.RepoManager
	UseCaseManager manager.UseCaseManager
}

func NewConfig() *Config {
	//dbHost := os.Getenv("DB_HOST")
	//dbPort := os.Getenv("DB_PORT")
	//dbName := os.Getenv("DB_NAME")
	//dbUser := os.Getenv("DB_USER")
	//dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := "localhost"
	dbPort := "5432"
	dbName := "song"
	dbUser := "postgres"
	dbPassword := "stauffenberg"

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	infraManager := manager.NewInfra(dataSourceName)
	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)

	config := new(Config)
	config.InfraManager = infraManager
	config.RepoManager = repoManager
	config.UseCaseManager = useCaseManager

	return config
}
