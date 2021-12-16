package process_transaction

type TransactionDtoInput struct {
	ID        string  `json:"id"`
	AccountID string  `json:"account_id"`
	Ammount   float64 `json:"ammount"`
}

type TransactionDtoOutput struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}
