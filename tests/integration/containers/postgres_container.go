package containers

import (
	"github.com/chinmayvivek/SwiftSchool/config"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

const (
	postgresImage    = "postgres:14.1-alpine"
	postgresPort     = 5432
	postgresUser     = "postgres"
	postgresPassword = "postgres"
	postgresDatabase = "postgres"
)

type PostgresContainer struct {
	container *testcontainers.Container
	Client    *postgres.PostgresContainer
}

func NewPostgresContainer() *PostgresContainer {
	return &PostgresContainer{}
}

func (p *PostgresContainer) InitContainer() (config.PostgresConfig, error) {

	pgConfig := config.PostgresConfig{
		Host:     "localhost",
		Port:     postgresPort,
		Username: postgresUser,
		Password: postgresPassword,
		Database: postgresDatabase,
	}

	return pgConfig, nil
}
