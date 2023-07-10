package types

import "github.com/metaphi-org/go-infra-sdk/sqs"

type InternalBankStatementJson struct {
	Transactions     []Transaction   `json:"transactions" validate:"required,dive"` // dive was to check each element.
	StatementDetails Details         `json:"details" validate:"required"`
	FraudChecks      sqs.FraudChecks `json:"fraudChecks"`
}

type Details struct {
	AccountNumber string `json:"accountNumber" validate:"required"`
	AccountName   string `json:"accountName" validate:"required"`
	BankName      string `json:"bankName" validate:"required"`
}

type Transaction struct {
	Id        int     `json:"id" validate:"required"`
	TxnDate   string  `json:"txnDate" validate:"required"`
	Narration string  `json:"narration" validate:"required"`
	Balance   float64 `json:"balance" validate:"required"`
	Amount    float64 `json:"amt" validate:"required"`
}
