package cmd

import (
	"money-transfer/internal/infrastructure/container"
	"money-transfer/internal/server"
)

func Run() {
	server.StartService(container.New())
}
