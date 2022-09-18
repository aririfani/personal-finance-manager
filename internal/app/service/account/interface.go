package account

import "context"

type Service interface {
	// CreateAccount ...
	CreateAccount(ctx context.Context, account Account) (returnData Account, err error)

	// UpdateAccount ...
	UpdateAccount(ctx context.Context, account Account) (returnData Account, err error)

	// GetAccountByID ...
	GetAccountByID(ctx context.Context, id int64) (returnData Account, err error)

	// GetAllAccount ...
	GetAllAccount(ctx context.Context, req GetAllAccountReq) (returnData GetAllAccountRes, err error)

	// DeleteAccount ...
	DeleteAccount(ctx context.Context, id int64) (returnData Account, err error)
}
