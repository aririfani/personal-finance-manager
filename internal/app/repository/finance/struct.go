package finance

import (
	"gorm.io/gorm"
	"time"
)

type Finance struct {
	ID               int64          `json:"id"  deepcopier:"ID"`
	Title            string         `json:"title" deepcopier:"Title"`
	FinanceAccountID int64          `json:"finance_account_id" deepcopier:"FinanceAccountID"`
	Amount           float64        `json:"amount" deepcopier:"Amount"`
	Description      string         `json:"description" deepcopier:"Description"`
	UserID           int64          `json:"user_id" deepcopier:"UserID"`
	Type             string         `json:"type" deepcipier:"type"`
	TransactionDate  string         `json:"transaction_date" deepcopier:"TransactionDate"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at"`
}

// TableName ...
func (Finance) TableName() string {
	return "finances"
}

// GetAllFinanceReq ...
type GetAllFinanceReq struct {
	UserID    int64  `json:"user_id"`
	Page      int    `json:"page"`
	Limit     int    `json:"limit"`
	Title     string `json:"title"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Type      string `json:"type"`
}

type Res struct {
	ID               int64     `json:"id"  deepcopier:"ID"`
	Title            string    `json:"title" deepcipier:"Title"`
	FinanceAccountID int64     `json:"finance_account_id" deepcopier:"FinanceAccountID"`
	Amount           float64   `json:"amount" deepcopier:"Amount"`
	Description      string    `json:"description" deepcopier:"Description"`
	UserID           int64     `json:"user_id" deepcopier:"UserID"`
	Type             string    `json:"type" deepcopier:"Type"`
	TransactionDate  time.Time `json:"transaction_date" deepcopier:"TransactionDate"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
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

// TotalTransactionMonthlyRes ...
type TotalTransactionMonthlyRes struct {
	Month    string  `json:"month" deepcopier:"Month"`
	MonthNum string  `json:"month_num" deepcopier:"MonthNum"`
	Year     string  `json:"year" deepcopier:"Year"`
	Total    float64 `json:"total" deepcopier:"Total"`
	Type     string  `json:"type" deepcopier:"Type"`
}
