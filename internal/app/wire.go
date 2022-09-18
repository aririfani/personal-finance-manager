package app

import (
	"github.com/aririfani/personal-finance-manager/internal/app/repository"
	accountRepo "github.com/aririfani/personal-finance-manager/internal/app/repository/account"
	financeRepo "github.com/aririfani/personal-finance-manager/internal/app/repository/finance"
	userRepo "github.com/aririfani/personal-finance-manager/internal/app/repository/user"
	"github.com/aririfani/personal-finance-manager/internal/app/service"
	accountService "github.com/aririfani/personal-finance-manager/internal/app/service/account"
	financeService "github.com/aririfani/personal-finance-manager/internal/app/service/finance"
	userService "github.com/aririfani/personal-finance-manager/internal/app/service/user"
	"github.com/aririfani/personal-finance-manager/internal/pkg/driver/driversql"
	"github.com/aririfani/personal-finance-manager/internal/pkg/token"
)

func WiringRepository(db *driversql.Database) *repository.Repositories {
	return &repository.Repositories{
		User:    userRepo.NewRepo(userRepo.NewDB(db)),
		Account: accountRepo.NewRepo(accountRepo.NewDB(db)),
		Finance: financeRepo.NewRepo(financeRepo.NewDB(db)),
	}
}

func WiringService(repo *repository.Repositories, token token.IToken) *service.Services {
	return &service.Services{
		User:    userService.NewSrv(repo, token),
		Account: accountService.NewSrv(repo),
		Finance: financeService.NewSrv(repo),
	}
}
