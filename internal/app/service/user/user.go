package user

import (
	"context"
	"github.com/aririfani/personal-finance-manager/internal/app/repository"
	"github.com/aririfani/personal-finance-manager/internal/app/repository/user"
	"github.com/ulule/deepcopier"
)

// srv ...
type srv struct {
	repo *repository.Repositories
}

// NewSrv ..
func NewSrv(repo *repository.Repositories) Service {
	return &srv{
		repo: repo,
	}
}

// CreateUser ...
func (s *srv) CreateUser(ctx context.Context, param User) (returnData User, err error) {
	userRepo := user.User{}
	_ = deepcopier.Copy(param).To(&userRepo)

	res, err := s.repo.User.CreateUser(ctx, userRepo)
	if err != nil {
		return
	}
	_ = deepcopier.Copy(res).To(&returnData)

	return
}
