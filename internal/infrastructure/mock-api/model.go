package mockapi

type (
	CheckValidAccountRequest struct {
		AccountNumber string `json:"account_number"`
		BankCode      string `json:"bank_code"`
	}

	TransferMoneyRequest struct {
		ExternalID  string  `json:"external_id,omitempty"`
		FromAccount string  `json:"from_account"`
		ToAccount   string  `json:"to_account"`
		ToBankCode  string  `json:"to_bank_code"`
		Amount      float64 `json:"amount"`
	}
)

type (
	CheckValidAccountResponse struct {
		Data    AccountDetail `json:"data"`
		Message string        `json:"message"`
		Status  int           `json:"status"`
	}

	AccountDetail struct {
		Name          string `json:"name"`
		AccountNumber string `json:"account_number"`
		BankCode      string `json:"bank_code"`
		Status        string `json:"status"`
	}

	TransferMoneyResponse struct {
		Data    Transaction `json:"data"`
		Message string      `json:"message"`
		Status  int         `json:"status"`
	}

	Transaction struct {
		TransactionDetail TransferMoneyRequest `json:"transaction_detail"`
		TransactionID     string               `json:"transaction_id"`
		ExternalID        string               `json:"external_id"`
		Status            string               `json:"status"`
		TransactionDate   string               `json:"transaction_date"`
	}
)
