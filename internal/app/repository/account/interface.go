package account

import "context"

type Repository interface {
	// CreateAccount ...
	CreateAccount(ctx context.Context, account Account) (returnData Account, err error)

	// UpdateAccount ...
	UpdateAccount(ctx context.Context, account Account) (returnData Account, err error)

	// GetAccountByID ...
	GetAccountByID(ctx context.Context, id int64) (returnData Account, err error)

	// GetAllAccount ...
	GetAllAccount(ctx context.Context, req GetAllAccountReq) (returnData []Account, err error)

	// DeleteAccount ...
	DeleteAccount(ctx context.Context) (returnData Account, err error)

	// CountTotalAccount ...
	CountTotalAccount(ctx context.Context, userID int64) (total int64, err error)
}
