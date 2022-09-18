package handler

import (
	"github.com/aririfani/personal-finance-manager/internal/app/handler/account"
	"github.com/aririfani/personal-finance-manager/internal/app/handler/finance"
	"github.com/aririfani/personal-finance-manager/internal/app/handler/user"
	"github.com/aririfani/personal-finance-manager/internal/app/service"
)

// Handler ...
type Handler interface {
	UserHandler() user.Handler
	AccountHandler() account.Handler
	FinanceHandler() finance.Handler
}

// handler ...
type handler struct {
	user    user.Handler
	account account.Handler
	finance finance.Handler
}

// NewHandler ...
func NewHandler(service *service.Services) Handler {
	h := new(handler)
	h.user = user.NewHandler(service)
	h.account = account.NewHandler(service)
	h.finance = finance.NewHandler(service)

	return h
}

// UserHandler ...
func (h *handler) UserHandler() user.Handler {
	return h.user
}

// AccountHandler ...
func (h *handler) AccountHandler() account.Handler {
	return h.account
}

// FinanceHandler ...
func (h *handler) FinanceHandler() finance.Handler {
	return h.finance
}
