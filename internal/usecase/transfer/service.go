package transfer

import (
	"context"
	"money-transfer/internal/domain/entities"
	"money-transfer/internal/domain/repositories"
)

type (
	TransferSvc interface {
		Create(ctx context.Context, transfer entities.Transfer) (res DefaultResponse, err error)
		UpdateStatus(ctx context.Context, paymentRef string, status string) (res DefaultResponse, err error)
		CheckValidAccount(ctx context.Context, req CheckValidAccount) (res DefaultResponse, err error)
	}

	transferSvc struct {
		transferRepo repositories.TransferRepo
	}
)

func NewTransferSvc(transferRepo repositories.TransferRepo) TransferSvc {
	if transferRepo == nil {
		panic("ProductRepo is required")
	}

	return &transferSvc{
		transferRepo: transferRepo,
	}
}

func (s *transferSvc) Create(ctx context.Context, transfer entities.Transfer) (res DefaultResponse, err error) {
	return
}

func (s *transferSvc) UpdateStatus(ctx context.Context, paymentRef string, status string) (res DefaultResponse, err error) {
	return
}

func (s *transferSvc) CheckValidAccount(ctx context.Context, req CheckValidAccount) (res DefaultResponse, err error) {
	return
}
