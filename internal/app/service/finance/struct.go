package finance

import (
	"github.com/aririfani/personal-finance-manager/internal/app/service/account"
	"time"
)

type Finance struct {
	ID              int64           `json:"id" deepcopier:"ID"`
	Title           string          `json:"title" deepcopier:"Title"`
	AccountID       int64           `json:"account_id" deepcopier:"AccountID"`
	Account         account.Account `json:"account"`
	Amount          float64         `json:"amount" deepcopier:"Amount"`
	Description     string          `json:"description" deepcopier:"Description"`
	UserID          int64           `json:"user_id" deepcopier:"UserID"`
	Type            string          `json:"type" deepcipier:"type"`
	TransactionDate string          `json:"transaction_date" deepcopier:"TransactionDate"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

// GetAllFinanceReq ...
type GetAllFinanceReq struct {
	UserID    int64  `json:"user_id" deepcopier:"UserID"`
	Page      int    `json:"page" deepcopier:"Page"`
	Limit     int    `json:"limit" deepcopier:"Limit"`
	Title     string `json:"title"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Type      string `json:"type"`
}

// GetAllFinanceRes ...
type GetAllFinanceRes struct {
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	Limit    int         `json:"limit"`
	LastPage int         `json:"last_page"`
	Data     interface{} `json:"data"`
}

// GetTotalTransaction ...
type GetTotalTransaction struct {
	UserID    int64  `json:"user_id" deepcopier:"UserID"`
	StartDate string `json:"start_date" deepcopier:"StartDate"`
	EndDate   string `json:"end_date" deepcopier:"EndDate"`
	Type      string `json:"type" deepcopier:"Type"`
}

// TotalTransactionRes ...
type TotalTransactionRes struct {
	TransactionDate time.Time `json:"transaction_date" deepcopier:"TransactionDate"`
	Total           float64   `json:"total" deepcopier:"Total"`
	Type            string    `json:"type" deepcopier:"Type"`
}
