package entities

import "time"

type (
	Transfer struct {
		ID            int       `db:"id" json:"id"`
		TrxID         string    `db:"trx_id" json:"trx_id"`
		PaymentRef    string    `db:"payment_ref" json:"payment_ref"`
		FromAccountID string    `db:"from_account_id" json:"from_account_id"`
		ToAccountID   string    `db:"to_account_id" json:"to_account_id"`
		Amount        float64   `db:"amount" json:"amount"`
		Status        string    `db:"status" json:"status"`
		CreatedAt     time.Time `db:"created_at" json:"created_at"`
		UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
	}
)
