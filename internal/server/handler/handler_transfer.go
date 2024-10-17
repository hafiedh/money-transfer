package handler

import (
	"log/slog"
	"money-transfer/internal/pkg"
	"money-transfer/internal/usecase/transfer"
	"net/http"

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
	ctx := e.Request().Context()

	var req transfer.CheckValidAccount

	if err = e.Bind(&req); err != nil {
		slog.ErrorContext(ctx, "CheckBankAccount: %v", err)
		return e.JSON(http.StatusBadRequest,
			pkg.DefaultResponse{
				Message: "Invalid request",
				Status:  http.StatusBadRequest,
				Data:    struct{}{},
			},
		)
	}

	if err := e.Validate(req); err != nil {
		slog.Error("Validation error: %v", err)
		return e.JSON(http.StatusBadRequest,
			pkg.DefaultResponse{
				Message: "Invalid request",
				Status:  http.StatusBadRequest,
				Data:    struct{}{},
			},
		)
	}

	res, err := h.transferService.CheckValidAccount(ctx, req)
	if err != nil {
		slog.ErrorContext(ctx, "CheckBankAccount: %v", err)
		return e.JSON(http.StatusInternalServerError,
			pkg.DefaultResponse{
				Message: "Internal server error",
				Status:  http.StatusInternalServerError,
				Data:    struct{}{},
			},
		)
	}

	return e.JSON(http.StatusOK, res)

}

func recoveryPanicHandler() {
	if r := recover(); r != nil {
		slog.Error("Panic Recovered: %v", r)
	}
}
