package handler

import (
	"money-transfer/internal/infrastructure/container"

	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo, cnt *container.Container) {
	h := SetupHandler(cnt).Validate()

	e.GET("/", h.healthCheckHandler.HealthCheck)
	banks := e.Group("/v1/banks")
	{
		banks.POST("/check-account", h.transferHandler.CheckBankAccount)
		banks.POST("/transfer", h.transferHandler.TransferMoney)
		banks.POST("/callback", h.transferHandler.TransferCallback)
	}
}
