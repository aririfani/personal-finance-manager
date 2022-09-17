package account

import "time"

type Account struct {
	ID          int64     `json:"id" deepcopier:"ID"`
	Name        string    `json:"name" deepcopier:"Name"`
	Type        string    `json:"type" deepcopier:"Type"`
	Description string    `json:"description" deepcopier:"Description"`
	UserID      int64     `json:"user_id" deepcopier:"UserID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// GetAllAccountReq ...
type GetAllAccountReq struct {
	UserID int64 `json:"user_id" deepcopier:"UserID"`
	Page   int   `json:"page" deepcopier:"Page"`
	Limit  int   `json:"limit" deepcopier:"Limit"`
}

// GetAllAccountRes ...
type GetAllAccountRes struct {
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	Limit    int         `json:"limit"`
	LastPage int         `json:"last_page"`
	Data     interface{} `json:"data"`
}
