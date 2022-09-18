package finance

import "context"

type Service interface {
	CreateFinance(ctx context.Context, finance Finance) (returnData Finance, err error)
	GetAllFinance(ctx context.Context, req GetAllFinanceReq) (returnData GetAllFinanceRes, err error)
	UpdateFinance(ctx context.Context, finance Finance) (returnData Finance, err error)
}
