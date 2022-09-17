package account

import "context"

type repo struct {
	DB Repository
}

func NewRepo(db Repository) Repository {
	return &repo{
		DB: db,
	}
}

// CreateAccount ...
func (r *repo) CreateAccount(ctx context.Context, account Account) (returnData Account, err error) {
	returnData, err = r.DB.CreateAccount(ctx, account)
	if err != nil {
		return
	}

	return
}

// UpdateAccount ...
func (r *repo) UpdateAccount(ctx context.Context, account Account) (returnData Account, err error) {
	returnData, err = r.DB.UpdateAccount(ctx, account)
	if err != nil {
		return
	}

	return
}

// GetAccountByID ...
func (r *repo) GetAccountByID(ctx context.Context, id int64) (returnData Account, err error) {
	returnData, err = r.DB.GetAccountByID(ctx, id)
	if err != nil {
		return
	}

	return
}

// GetAllAccount ...
func (r *repo) GetAllAccount(ctx context.Context, req GetAllAccountReq) (returnData []Account, err error) {
	returnData, err = r.DB.GetAllAccount(ctx, req)
	if err != nil {
		return
	}

	return
}

// DeleteAccount ...
func (r *repo) DeleteAccount(ctx context.Context) (returnData Account, err error) {
	returnData, err = r.DB.DeleteAccount(ctx)
	if err != nil {
		return
	}

	return
}

// CountTotalAccount ...
func (r *repo) CountTotalAccount(ctx context.Context, userID int64) (total int64, err error) {
	total, err = r.DB.CountTotalAccount(ctx, userID)
	if err != nil {
		return
	}

	return
}
