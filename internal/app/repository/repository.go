package repository

import (
	"github.com/aririfani/personal-finance-manager/internal/app/repository/account"
	"github.com/aririfani/personal-finance-manager/internal/app/repository/user"
)

// Repositories ...
type Repositories struct {
	User    user.Repository
	Account account.Repository
}
