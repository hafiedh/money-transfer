package handler

import (
	"log/slog"
	"money-transfer/internal/usecase/transfer"

	"github.com/labstack/echo/v4"
)

type (
	transferHandler struct {
		transferService transfer.TransferSvc
	}
)

func NewTransferHandler() *transferHandler {
	return &transferHandler{}
}

func (h *transferHandler) Set(svc transfer.TransferSvc) *transferHandler {
	h.transferService = svc
	return h
}

func (h *transferHandler) Validate() *transferHandler {
	if h.transferService == nil {
		panic("transferService is nil")
	}

	return h
}

func (h *transferHandler) TransferMoney(e echo.Context) (err error) {
	defer recoveryPanicHandler()
	return
}

func (h *transferHandler) TransferCallback(e echo.Context) (err error) {
	defer recoveryPanicHandler()
	return
}

func (h *transferHandler) CheckBankAccount(e echo.Context) (err error) {
	defer recoveryPanicHandler()
	return
}

func recoveryPanicHandler() {
	if r := recover(); r != nil {
		slog.Error("Panic Recovered: %v", r)
	}
}
