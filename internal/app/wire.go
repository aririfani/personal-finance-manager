package app

import (
	"github.com/aririfani/personal-finance-manager/internal/app/repository"
	userRepo "github.com/aririfani/personal-finance-manager/internal/app/repository/user"
	"github.com/aririfani/personal-finance-manager/internal/app/service"
	userService "github.com/aririfani/personal-finance-manager/internal/app/service/user"
	"github.com/aririfani/personal-finance-manager/internal/pkg/driver/driversql"
)

func WiringRepository(db *driversql.Database) *repository.Repositories {
	return &repository.Repositories{
		User: userRepo.NewRepo(userRepo.NewDB(db)),
	}
}

func WiringService(repo *repository.Repositories) *service.Services {
	return &service.Services{
		User: userService.NewSrv(repo),
	}
}
