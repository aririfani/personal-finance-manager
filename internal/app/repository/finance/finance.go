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
