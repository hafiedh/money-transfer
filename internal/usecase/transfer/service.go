package transfer

import (
	"context"
	"fmt"
	"log/slog"
	"money-transfer/internal/domain/entities"
	"money-transfer/internal/domain/repositories"
	mockapi "money-transfer/internal/infrastructure/mock-api"
	"money-transfer/internal/pkg"
)

type (
	TransferSvc interface {
		Create(ctx context.Context, transfer entities.Transfer) (res pkg.DefaultResponse, err error)
		UpdateStatus(ctx context.Context, paymentRef string, status string) (res pkg.DefaultResponse, err error)
		CheckValidAccount(ctx context.Context, req CheckValidAccount) (res pkg.DefaultResponse, err error)
	}

	transferSvc struct {
		transferRepo repositories.TransferRepo
		mockApi      mockapi.Bank
	}
)

func NewTransferSvc(transferRepo repositories.TransferRepo, mock mockapi.Bank) TransferSvc {
	if transferRepo == nil {
		panic("ProductRepo is required")
	}
	if mock == nil {
		panic("MockApi is required")
	}

	return &transferSvc{
		transferRepo: transferRepo,
		mockApi:      mock,
	}
}

func (s *transferSvc) Create(ctx context.Context, transfer entities.Transfer) (res pkg.DefaultResponse, err error) {
	return
}

func (s *transferSvc) UpdateStatus(ctx context.Context, paymentRef string, status string) (res pkg.DefaultResponse, err error) {
	return
}

func (s *transferSvc) CheckValidAccount(ctx context.Context, req CheckValidAccount) (res pkg.DefaultResponse, err error) {
	wrapperReq := mockapi.CheckValidAccountRequest{
		AccountNumber: req.AccountNumber,
		BankCode:      req.BankCode,
	}
	wrapperRes, err := s.mockApi.CheckValidAccount(ctx, wrapperReq)
	if err != nil {
		slog.ErrorContext(ctx, "[service] CheckValidAccount: %v", err)
		err = fmt.Errorf("cannot check valid account")
		return
	}
	res = pkg.DefaultResponse{
		Message: "Success",
		Status:  200,
		Data:    wrapperRes.Data,
	}

	return

}
