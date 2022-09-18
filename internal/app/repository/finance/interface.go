package finance

import "context"

type Repository interface {
	// CreateFinance ...
	CreateFinance(ctx context.Context, finance Finance) (returnData Finance, err error)

	// GetAllFinance ...
	GetAllFinance(ctx context.Context, param GetAllFinanceReq) (returnData []Finance, err error)

	// CountTotalFinance ...
	CountTotalFinance(ctx context.Context, req GetAllFinanceReq) (total int64, err error)

	// UpdateFinance ...
	UpdateFinance(ctx context.Context, finance Finance) (returnData Finance, err error)

	// GetFinanceByID ...
	GetFinanceByID(ctx context.Context, id int64) (returnData Finance, err error)
}
