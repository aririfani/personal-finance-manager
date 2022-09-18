package finance

import (
	"context"
	"github.com/aririfani/personal-finance-manager/internal/pkg/driver/driversql"
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

	err = query.Find(&returnData).Error

	if err != nil {
		return
	}

	return
}

// CountTotalFinance ...
func (d *db) CountTotalFinance(ctx context.Context, userID int64) (total int64, err error) {
	err = d.DB.Instance.WithContext(ctx).Where("user_id =?", userID).Table("finances").Count(&total).Error
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
