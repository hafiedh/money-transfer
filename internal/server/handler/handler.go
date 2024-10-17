package handler

import (
	"money-transfer/internal/infrastructure/container"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Handler struct {
	postgres           *pgxpool.Pool
	healthCheckHandler *healthCheckHandler
	transferHandler    *transferHandler
}

func SetupHandler(container *container.Container) *Handler {
	return &Handler{
		postgres:           container.PostgresDB,
		healthCheckHandler: NewHealthCheckHandler().SetHealthCheckService(container.HealthCheckService).Validate(),
		transferHandler:    NewTransferHandler().Set(container.TransferService).Validate(),
	}
}

func (h *Handler) Validate() *Handler {
	if h.healthCheckHandler == nil {
		panic("healthCheckHandler is nil")
	}
	if h.transferHandler == nil {
		panic("transferHandler is nil")
	}
	return h
}
