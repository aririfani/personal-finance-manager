package account

import (
	"context"
	"github.com/aririfani/personal-finance-manager/internal/app/repository"
	"github.com/aririfani/personal-finance-manager/internal/app/repository/account"
	"github.com/ulule/deepcopier"
	"math"
)

type srv struct {
	repo *repository.Repositories
}

// NewSrv ..
func NewSrv(repo *repository.Repositories) Service {
	return &srv{
		repo: repo,
	}
}

// CreateAccount ...
func (s srv) CreateAccount(ctx context.Context, param Account) (returnData Account, err error) {
	accountRepo := account.Account{}
	_ = deepcopier.Copy(param).To(&accountRepo)

	res, err := s.repo.Account.CreateAccount(ctx, accountRepo)
	if err != nil {
		return
	}

	_ = deepcopier.Copy(res).To(&returnData)

	return
}

// UpdateAccount ...
func (s srv) UpdateAccount(ctx context.Context, param Account) (returnData Account, err error) {
	accountRepo := account.Account{}
	_ = deepcopier.Copy(param).To(&accountRepo)
	res, err := s.repo.Account.UpdateAccount(ctx, accountRepo)
	if err != nil {
		return
	}

	_ = deepcopier.Copy(res).To(&returnData)

	return
}

// GetAccountByID ...
func (s srv) GetAccountByID(ctx context.Context, id int64) (returnData Account, err error) {
	res, err := s.repo.Account.GetAccountByID(ctx, id)
	if err != nil {
		return
	}
	_ = deepcopier.Copy(res).To(&returnData)
	return
}

// GetAllAccount ...
func (s srv) GetAllAccount(ctx context.Context, req GetAllAccountReq) (returnData GetAllAccountRes, err error) {
	var accountRepo []account.Res
	var accountReq account.GetAllAccountReq
	_ = deepcopier.Copy(req).To(&accountReq)
	res, err := s.repo.Account.GetAllAccount(ctx, accountReq)
	if err != nil {
		return
	}

	for _, v := range res {
		var accountTmp account.Res
		_ = deepcopier.Copy(v).To(&accountTmp)
		accountRepo = append(accountRepo, accountTmp)
	}

	totalData, err := s.repo.Account.CountTotalAccount(ctx, accountReq.UserID)
	if err != nil {
		return
	}

	if len(accountRepo) == 0 {
		accountRepo = []account.Res{}
	}

	returnData.Data = accountRepo
	returnData.Total = totalData
	returnData.Page = 1
	returnData.Limit = 10

	if req != (GetAllAccountReq{}) {
		if req.Limit != 0 {
			returnData.Limit = req.Limit
		}

		if req.Page != 0 {
			returnData.Page = req.Page
		}

		returnData.LastPage = int(math.Ceil(float64(returnData.Total) / float64(returnData.Limit)))
	}

	return
}

// DeleteAccount ...
func (s srv) DeleteAccount(ctx context.Context) (returnData Account, err error) {
	return
}
