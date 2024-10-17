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
		UpdateStatus(ctx context.Context, paymentRef, status string) error
	}

	productRepo struct {
		db *pgxpool.Pool
	}
)

func NewTransferRepo(db *pgxpool.Pool) TransferRepo {
	if db == nil {
		panic("ProductRepo: db is nil")
	}

	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(ctx context.Context, transfer entities.Transfer) error {
	_, err := r.db.Exec(ctx, "INSERT INTO transfers (from_account_id, to_account_id, amount, status) VALUES ($1, $2, $3, $4)",
		transfer.FromAccountID, transfer.ToAccountID, transfer.Amount, transfer.Status)
	if err != nil {
		slog.ErrorContext(ctx, "[ProductRepo.Create] Error: %v", err)
		err = fmt.Errorf("failed to create transfer")
		return err
	}

	return nil
}

func (r *productRepo) UpdateStatus(ctx context.Context, paymentRef, status string) error {
	_, err := r.db.Exec(ctx, "UPDATE transfers SET status = $1 WHERE payment_ref = $2", status, paymentRef)
	if err != nil {
		slog.ErrorContext(ctx, "[ProductRepo.UpdateStatus] Error: %v", err)
		err = fmt.Errorf("failed to update transfer status")
		return err
	}

	return nil
}
