package mockapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"money-transfer/internal/config"
	"net/http"
)

type (
	Bank interface {
		TransferMoney(ctx context.Context, req TransferMoneyRequest) (resp TransferMoneyResponse, err error)
		CheckValidAccount(ctx context.Context, req CheckValidAccountRequest) (resp CheckValidAccountResponse, err error)
	}

	bank struct{}
)

func NewBank() Bank {
	return &bank{}
}

func (b *bank) TransferMoney(ctx context.Context, req TransferMoneyRequest) (resp TransferMoneyResponse, err error) {
	payload, err := json.Marshal(req)
	if err != nil {
		slog.ErrorContext(ctx, "[Bank.TransferMoney]", "error when marshal request", err.Error())
		err = fmt.Errorf("cannot transfer money")
		return
	}
	request, err := http.NewRequest(http.MethodPost, config.GetString("postman.mocks.url"), bytes.NewBuffer(payload))
	if err != nil {
		slog.ErrorContext(ctx, "[Bank.TransferMoney]", "error when create request", err.Error())
		err = fmt.Errorf("cannot transfer money")
		return
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		slog.ErrorContext(ctx, "[Bank.TransferMoney]", "error when do request", err.Error())
		err = fmt.Errorf("cannot transfer money")
		return
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		slog.ErrorContext(ctx, "[Bank.TransferMoney]", "error when decode response", err.Error())
		err = fmt.Errorf("cannot transfer money")
		return
	}

	return

}

func (b *bank) CheckValidAccount(ctx context.Context, req CheckValidAccountRequest) (resp CheckValidAccountResponse, err error) {
	payload, err := json.Marshal(req)
	if err != nil {
		slog.ErrorContext(ctx, "[Bank.CheckValidAccount]", "error when marshal request", err.Error())
		err = fmt.Errorf("cannot check valid account")
		return
	}
	request, err := http.NewRequest(http.MethodPost, config.GetString("postman.mocks.url")+config.GetString("postman.mocks.checkAccount"), bytes.NewBuffer(payload))
	if err != nil {
		slog.ErrorContext(ctx, "[Bank.CheckValidAccount]", "error when create request", err.Error())
		err = fmt.Errorf("cannot check valid account")
		return
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		slog.ErrorContext(ctx, "[Bank.CheckValidAccount]", "error when do request", err.Error())
		err = fmt.Errorf("cannot check valid account")
		return
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		slog.ErrorContext(ctx, "[Bank.CheckValidAccount]", "error when decode response", err.Error())
		err = fmt.Errorf("cannot check valid account")
		return
	}

	return
}
