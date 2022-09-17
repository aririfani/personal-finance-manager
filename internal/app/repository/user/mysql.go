package user

import (
	"context"
	"github.com/aririfani/personal-finance-manager/internal/pkg/driver/driversql"
)

type db struct {
	DB *driversql.Database
}

// NewDB ...
func NewDB(d *driversql.Database) Repository {
	return &db{
		DB: d,
	}
}

// CreateUser ...
func (d *db) CreateUser(ctx context.Context, user User) (returnData User, err error) {
	err = d.DB.Instance.WithContext(ctx).Create(&user).Error
	if err != nil {
		return
	}

	returnData = User{
		ID:        user.ID,
		Email:     user.Email,
		FullName:  user.FullName,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return
}
