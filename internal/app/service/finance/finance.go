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

	totalData, err := s.repo.Finance.CountTotalFinance(ctx, financeReq)
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

// GetFinanceByID ...
func (s *srv) GetFinanceByID(ctx context.Context, id int64) (returnData Finance, err error) {
	res, err := s.repo.Finance.GetFinanceByID(ctx, id)
	if err != nil {
		return
	}

	_ = deepcopier.Copy(res).To(&returnData)
	return
}

// DeleteFinanceByID ...
func (s *srv) DeleteFinanceByID(ctx context.Context, id int64) (returnData Finance, err error) {
	res, err := s.repo.Finance.GetFinanceByID(ctx, id)
	if err != nil {
		return
	}

	err = s.repo.Finance.DeleteFinanceByID(ctx, id)
	if err != nil {
		return
	}

	_ = deepcopier.Copy(res).To(&returnData)
	return
}

// GetTotalTransactionDaily ...
func (s *srv) GetTotalTransactionDaily(ctx context.Context, req GetTotalTransaction) (returnData interface{}, err error) {
	financeRepo := finance.GetTotalTransaction{}
	_ = deepcopier.Copy(req).To(&financeRepo)

	res, err := s.repo.Finance.GetTotalTransactionDaily(ctx, financeRepo)

	returnData = res
	return
}

// GetTotalTransactionMonthly ...
func (s *srv) GetTotalTransactionMonthly(ctx context.Context, req GetTotalTransaction) (returnData interface{}, err error) {
	financeRepo := finance.GetTotalTransaction{}
	_ = deepcopier.Copy(req).To(&financeRepo)

	res, err := s.repo.Finance.GetTotalTransactionMonthly(ctx, financeRepo)

	returnData = res

	return
}
