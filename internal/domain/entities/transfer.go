package entities

import "time"

type (
	Transfer struct {
		ID            int       `db:"id"`
		TrxID         string    `db:"trx_id"`
		PaymentRef    string    `db:"payment_ref"`
		FromAccountID int       `db:"from_account_id"`
		ToAccountID   int       `db:"to_account_id"`
		Amount        float64   `db:"amount"`
		Status        string    `db:"status"`
		CreatedAt     time.Time `db:"created_at"`
		UpdatedAt     time.Time `db:"updated_at"`
	}
)
