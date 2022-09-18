package finance

import "context"

type repo struct {
	DB Repository
}

func NewRepo(db Repository) Repository {
	return &repo{
		DB: db,
	}
}

// CreateFinance ...
func (r *repo) CreateFinance(ctx context.Context, finance Finance) (returnData Finance, err error) {
	returnData, err = r.DB.CreateFinance(ctx, finance)
	if err != nil {
		return
	}

	return
}

// GetAllFinance ...
func (r *repo) GetAllFinance(ctx context.Context, req GetAllFinanceReq) (returnData []Finance, err error) {
	returnData, err = r.DB.GetAllFinance(ctx, req)
	if err != nil {
		return
	}

	return
}

// CountTotalFinance ...
func (r *repo) CountTotalFinance(ctx context.Context, req GetAllFinanceReq) (total int64, err error) {
	total, err = r.DB.CountTotalFinance(ctx, req)
	if err != nil {
		return
	}

	return
}

// UpdateFinance ...
func (r *repo) UpdateFinance(ctx context.Context, finance Finance) (returnData Finance, err error) {
	returnData, err = r.DB.UpdateFinance(ctx, finance)
	if err != nil {
		return
	}

	return
}

// GetFinanceByID ...
func (r *repo) GetFinanceByID(ctx context.Context, id int64) (returnData Finance, err error) {
	returnData, err = r.DB.GetFinanceByID(ctx, id)
	if err != nil {
		return
	}

	return
}

// DeleteFinanceByID ...
func (r *repo) DeleteFinanceByID(ctx context.Context, id int64) (err error) {
	err = r.DB.DeleteFinanceByID(ctx, id)
	return
}

// GetTotalTransactionDaily ...
func (r *repo) GetTotalTransactionDaily(ctx context.Context, req GetTotalTransaction) (returnData []TotalTransactionRes, err error) {
	returnData, err = r.DB.GetTotalTransactionDaily(ctx, req)
	if err != nil {
		return
	}
	return
}

// GetTotalTransactionMonthly ...
func (r *repo) GetTotalTransactionMonthly(ctx context.Context, req GetTotalTransaction) (returnData []TotalTransactionMonthlyRes, err error) {
	returnData, err = r.DB.GetTotalTransactionMonthly(ctx, req)
	if err != nil {
		return
	}

	return
}
