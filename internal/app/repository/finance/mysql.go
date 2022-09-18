package finance

import (
	"context"
	"github.com/aririfani/personal-finance-manager/internal/pkg/driver/driversql"
	"strings"
)

type db struct {
	DB *driversql.Database
}

func NewDB(d *driversql.Database) Repository {
	return &db{
		DB: d,
	}
}

// CreateFinance ...
func (d *db) CreateFinance(ctx context.Context, finance Finance) (returnData Finance, err error) {
	err = d.DB.Instance.WithContext(ctx).Create(&finance).Error
	if err != nil {
		return
	}

	returnData = Finance{
		ID:               finance.ID,
		Title:            finance.Title,
		Description:      finance.Description,
		Amount:           finance.Amount,
		UserID:           finance.UserID,
		FinanceAccountID: finance.FinanceAccountID,
		CreatedAt:        finance.CreatedAt,
		UpdatedAt:        finance.UpdatedAt,
	}

	return
}

// GetAllFinance ...
func (d *db) GetAllFinance(ctx context.Context, req GetAllFinanceReq) (returnData []Finance, err error) {
	query := d.DB.Instance.WithContext(ctx).Table("finances").Where("user_id =?", req.UserID)
	if req != (GetAllFinanceReq{}) {
		offset := req.Limit * (req.Page - 1)
		query = query.Offset(offset).Limit(req.Limit)
	}

	if strings.Compare(req.Title, "") != 0 {
		searchTitle := strings.ToLower(req.Title)
		query = query.Where("LOWER(title) LIKE ?", "%"+searchTitle+"%")
	}

	if strings.Compare(req.Type, "") != 0 {
		query = query.Where("type =?", req.Type)
	}

	if strings.Compare(req.StartDate, "") != 0 && strings.Compare(req.EndDate, "") != 0 {
		query = query.Where("transaction_date >= ?", req.StartDate).Where("transaction_date < ?", req.EndDate)
	}

	err = query.Find(&returnData).Error

	if err != nil {
		return
	}

	return
}

// CountTotalFinance ...
func (d *db) CountTotalFinance(ctx context.Context, req GetAllFinanceReq) (total int64, err error) {
	query := d.DB.Instance.WithContext(ctx).Where("user_id =?", req.UserID).Table("finances")

	if strings.Compare(req.Title, "") != 0 {
		searchTitle := strings.ToLower(req.Title)
		query = query.Where("LOWER(title) LIKE ?", "%"+searchTitle+"%")
	}

	if strings.Compare(req.Type, "") != 0 {
		query = query.Where("type =?", req.Type)
	}

	if strings.Compare(req.StartDate, "") != 0 && strings.Compare(req.EndDate, "") != 0 {
		query = query.Where("transaction_date >= ?", req.StartDate).Where("transaction_date < ?", req.EndDate)
	}

	err = query.Count(&total).Error
	if err != nil {
		return
	}
	return
}

// UpdateFinance ...
func (d *db) UpdateFinance(ctx context.Context, finance Finance) (returnData Finance, err error) {
	err = d.DB.Instance.WithContext(ctx).Where("id =?", finance.ID).Updates(&finance).Error
	if err != nil {
		return
	}

	return
}

// GetFinanceByID ...
func (d *db) GetFinanceByID(ctx context.Context, id int64) (returnData Finance, err error) {
	err = d.DB.Instance.WithContext(ctx).Where("id =?", id).First(&returnData).Error
	if err != nil {
		return
	}

	return
}
