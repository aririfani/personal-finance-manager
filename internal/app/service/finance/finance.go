package finance

import (
	"context"
	"github.com/aririfani/personal-finance-manager/internal/app/repository"
	"github.com/aririfani/personal-finance-manager/internal/app/repository/finance"
	"github.com/ulule/deepcopier"
	"math"
)

type srv struct {
	repo *repository.Repositories
}

func NewSrv(repo *repository.Repositories) Service {
	return &srv{
		repo: repo,
	}
}

// CreateFinance ...
func (s *srv) CreateFinance(ctx context.Context, param Finance) (returnData Finance, err error) {
	financeRepo := finance.Finance{}
	_ = deepcopier.Copy(param).To(&financeRepo)

	res, err := s.repo.Finance.CreateFinance(ctx, financeRepo)
	if err != nil {
		return
	}

	_ = deepcopier.Copy(res).To(&returnData)
	return
}

// GetAllFinance ...
func (s *srv) GetAllFinance(ctx context.Context, req GetAllFinanceReq) (returnData GetAllFinanceRes, err error) {
	var financeRepo []finance.Res
	var financeReq finance.GetAllFinanceReq
	_ = deepcopier.Copy(req).To(&financeReq)
	res, err := s.repo.Finance.GetAllFinance(ctx, financeReq)
	if err != nil {
		return
	}

	for _, v := range res {
		var financeTmp finance.Res
		_ = deepcopier.Copy(v).To(&financeTmp)
		financeRepo = append(financeRepo, financeTmp)
	}

	totalData, err := s.repo.Finance.CountTotalFinance(ctx, financeReq.UserID)
	if err != nil {
		return
	}

	if len(financeRepo) == 0 {
		financeRepo = []finance.Res{}
	}

	returnData.Data = financeRepo
	returnData.Total = totalData
	returnData.Page = 1
	returnData.Limit = 10

	if req != (GetAllFinanceReq{}) {
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

// UpdateFinance ...
func (s *srv) UpdateFinance(ctx context.Context, param Finance) (returnData Finance, err error) {
	financeRepo := finance.Finance{}
	_ = deepcopier.Copy(param).To(&financeRepo)

	res, err := s.repo.Finance.UpdateFinance(ctx, financeRepo)
	if err != nil {
		return
	}

	_ = deepcopier.Copy(res).To(&returnData)
	return
}
