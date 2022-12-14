package finance

import "context"

type Service interface {
	CreateFinance(ctx context.Context, finance Finance) (returnData Finance, err error)
	GetAllFinance(ctx context.Context, req GetAllFinanceReq) (returnData GetAllFinanceRes, err error)
	UpdateFinance(ctx context.Context, finance Finance) (returnData Finance, err error)
	GetFinanceByID(ctx context.Context, id int64) (returnData Finance, err error)
	DeleteFinanceByID(ctx context.Context, id int64) (returnData Finance, err error)
	GetTotalTransactionDaily(ctx context.Context, req GetTotalTransaction) (returnData interface{}, err error)
	GetTotalTransactionMonthly(ctx context.Context, req GetTotalTransaction) (returnData interface{}, err error)
}
