package repositories

import (
	"context"
	"fmt"
	"log/slog"
	"money-transfer/internal/domain/entities"

	"github.com/jackc/pgx/v4/pgxpool"
)

type (
	TransferRepo interface {
		Create(ctx context.Context, transfer entities.Transfer) error
		UpdateStatus(ctx context.Context, transfer entities.Transfer) error
	}

	transferRepo struct {
		db *pgxpool.Pool
	}
)

func NewTransferRepo(db *pgxpool.Pool) TransferRepo {
	if db == nil {
		panic("ProductRepo: db is nil")
	}

	return &transferRepo{
		db: db,
	}
}

func (r *transferRepo) Create(ctx context.Context, transfer entities.Transfer) error {
	_, err := r.db.Exec(ctx, `INSERT INTO transfers (payment_ref, from_account_number, to_account_number, amount, status, trx_id)
		VALUES ($1, $2, $3, $4, $5, $6)`, transfer.PaymentRef, transfer.FromAccountID, transfer.ToAccountID, transfer.Amount, transfer.Status, transfer.TrxID)
	if err != nil {
		slog.ErrorContext(ctx, "[ProductRepo.Create] Error: %v", err)
		err = fmt.Errorf("failed to create transfer")
		return err
	}

	return nil

}

func (r *transferRepo) UpdateStatus(ctx context.Context, transfer entities.Transfer) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "[ProductRepo.UpdateStatus] Error: %v", err)
		err = fmt.Errorf("failed to update transfer status")
		return err
	}
	var amount float64
	var status string
	err = tx.QueryRow(ctx, `SELECT amount, status FROM transfers WHERE payment_ref = $1`, transfer.PaymentRef).Scan(&amount, &status)
	if err != nil {
		slog.ErrorContext(ctx, "[ProductRepo.UpdateStatus] Error: %v", err)
		err = fmt.Errorf("failed to update transfer status")
		return err
	}

	if amount != transfer.Amount || status != "pending" {
		slog.ErrorContext(ctx, "[ProductRepo.UpdateStatus] Error: %v", err)
		err = fmt.Errorf("invalid amount or status")
		return err
	}

	_, err = tx.Exec(ctx, `UPDATE transfers SET status = $1 WHERE payment_ref = $2`, transfer.Status, transfer.PaymentRef)
	if err != nil {
		slog.ErrorContext(ctx, "[ProductRepo.UpdateStatus] Error: %v", err)
		err = fmt.Errorf("failed to update transfer status")
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "[ProductRepo.UpdateStatus] Error: %v", err)
		err = fmt.Errorf("failed to update transfer status")
		return err
	}

	return nil
}
