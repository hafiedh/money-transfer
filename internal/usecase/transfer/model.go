package transfer

type (
	MoneyTransfer struct {
		ExternalID string  `json:"external_id" validate:"required"`
		Amount     float64 `json:"amount" validate:"required"`
		FromAcount string  `json:"from_acount" validate:"required"`
		ToAccount  string  `json:"to_account" validate:"required"`
		ToBankCode string  `json:"to_bank_code" validate:"required"`
	}

	CheckValidAccount struct {
		AccountNumber string `json:"account_number" validate:"required"`
		BankCode      string `json:"bank_code" validate:"required"`
	}

	TransferCallback struct {
		ExternalID string `json:"external_id" validate:"required"`
		Amount     string `json:"amount" validate:"required"`
		Status     string `json:"status" validate:"required"`
	}
)
