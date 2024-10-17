package container

import (
	"fmt"

	"money-transfer/internal/config"
	"money-transfer/internal/domain/repositories"
	mockapi "money-transfer/internal/infrastructure/mock-api"
	"money-transfer/internal/infrastructure/postgres"
	"money-transfer/internal/usecase/healthcheck"
	"money-transfer/internal/usecase/transfer"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Container struct {
	Config             *config.DefaultConfig
	PostgresDB         *pgxpool.Pool
	HealthCheckService healthcheck.Service
	TransferService    transfer.TransferSvc
}

func (c *Container) Validate() *Container {
	if c.Config == nil {
		panic("Config is nil")
	}
	if c.HealthCheckService == nil {
		panic("HealthCheckService is nil")
	}
	if c.PostgresDB == nil {
		panic("PostgresDB is nil")
	}
	if c.TransferService == nil {
		panic("TransferService is nil")
	}
	return c
}

func New() *Container {
	err := config.Load("env", ".env")
	if err != nil {
		fmt.Printf("Error loading config: %v", err)
		panic(err)
	}

	defConfig := &config.DefaultConfig{
		Apps: config.Apps{
			Name:     config.GetString("app.name"),
			Address:  config.GetString("address"),
			HttpPort: config.GetString("port"),
		},
	}

	postgresConfig := &config.PostgreSQLDB{
		Host:         config.GetString("postgresql.money_transfer_db.host"),
		User:         config.GetString("postgresql.money_transfer_db.user"),
		Password:     config.GetString("postgresql.money_transfer_db.password"),
		Name:         config.GetString("postgresql.money_transfer_db.db"),
		Port:         config.GetInt("postgresql.money_transfer_db.port"),
		SSLMode:      config.GetString("postgresql.money_transfer_db.ssl"),
		Schema:       config.GetString("postgresql.money_transfer_db.schema"),
		Debug:        config.GetBool("postgresql.money_transfer_db.debug"),
		PoolMaxConns: config.GetInt("postgresql.money_transfer_db.poolMaxConns"),
	}

	// repo and db
	postgresDB, err := postgres.NewDB(*postgresConfig)
	if err != nil {
		fmt.Printf("Error connecting to PostgreSQL database: %v", err)
	}
	transferRepo := repositories.NewTransferRepo(postgresDB)
	bankWrapper := mockapi.NewBank()

	// service
	healthCheckService := healthcheck.NewService().Validate()
	transferService := transfer.NewTransferSvc(transferRepo, bankWrapper)

	container := &Container{
		Config:             defConfig,
		HealthCheckService: healthCheckService,
		PostgresDB:         postgresDB,
		TransferService:    transferService,
	}
	container.Validate()
	return container

}
