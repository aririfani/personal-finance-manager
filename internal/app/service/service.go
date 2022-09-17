package service

import (
	"github.com/aririfani/personal-finance-manager/internal/app/service/account"
	"github.com/aririfani/personal-finance-manager/internal/app/service/user"
)

type Services struct {
	User    user.Service
	Account account.Service
}
