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
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
}

func (Account) TableName() string {
	return "accounts"
}

// GetAllAccountReq ...
type GetAllAccountReq struct {
	UserID int64 `json:"user_id"`
	Page   int   `json:"page"`
	Limit  int   `json:"limit"`
}

type Res struct {
	ID          int64     `json:"id" deepcopier:"ID"`
	Name        string    `json:"name" deepcopier:"Name"`
	Type        string    `json:"type" deepcopier:"Type"`
	Description string    `json:"description" deepcopier:"Description"`
	UserID      int64     `json:"user_id" deepcopier:"UserID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
