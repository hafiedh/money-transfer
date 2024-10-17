package transfer

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log/slog"
	"money-transfer/internal/domain/entities"
	"money-transfer/internal/domain/repositories"
	mockapi "money-transfer/internal/infrastructure/mock-api"
	"money-transfer/internal/pkg"
	"strings"
	"time"
)

type (
	TransferSvc interface {
		Create(ctx context.Context, req MoneyTransfer) (res pkg.DefaultResponse, err error)
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

func (s *transferSvc) Create(ctx context.Context, req MoneyTransfer) (res pkg.DefaultResponse, err error) {
	paymentRef, err := generatePaymentRef(req.FromAccount)
	if err != nil {
		slog.ErrorContext(ctx, "[service][generatePaymentRef]: %v", err)
		err = fmt.Errorf("cannot generate payment ref")
		return
	}

	transferWrapper := mockapi.TransferMoneyRequest{
		ExternalID:  paymentRef,
		FromAccount: req.FromAccount,
		ToAccount:   req.ToAccount,
		ToBankCode:  req.ToBankCode,
		Amount:      req.Amount,
	}

	resWrapper, errWrapper := s.mockApi.TransferMoney(ctx, transferWrapper)
	if errWrapper != nil {
		slog.ErrorContext(ctx, "[service][wrapper] Create: %v", err)
		err = errWrapper
		return
	}

	entryData := entities.Transfer{
		PaymentRef:    paymentRef,
		TrxID:         resWrapper.Data.TransactionID,
		FromAccountID: req.FromAccount,
		ToAccountID:   req.ToAccount,
		Amount:        req.Amount,
		Status:        "pending",
	}

	err = s.transferRepo.Create(ctx, entryData)
	if err != nil {
		slog.ErrorContext(ctx, "[service][repo] Create: %v", err)
		err = fmt.Errorf("cannot create transfer")
		return
	}

	res = pkg.DefaultResponse{
		Message: "Success",
		Status:  201,
		Data:    struct{}{},
	}

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

func generatePaymentRef(AccountNumber string) (string, error) {
	randomBytes := make([]byte, 5)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	randomString := strings.ToUpper(hex.EncodeToString(randomBytes))
	phoneLast5 := AccountNumber[len(AccountNumber)-5:]
	secondsInDay := time.Now().Hour()*3600 + time.Now().Minute()*60 + time.Now().Second()
	timeComponent := fmt.Sprintf("%05d", secondsInDay)
	paymentCode := fmt.Sprintf("TF-%s%s%s", randomString[:3], phoneLast5[:2], timeComponent[:5])

	return paymentCode, nil
}
