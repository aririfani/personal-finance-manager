package handler

import (
	"github.com/aririfani/personal-finance-manager/internal/app/handler/user"
	"github.com/aririfani/personal-finance-manager/internal/app/service"
)

// Handler ...
type Handler interface {
	UserHandler() user.Handler
}

// handler ...
type handler struct {
	user user.Handler
}

// NewHandler ...
func NewHandler(service *service.Services) Handler {
	h := new(handler)
	h.user = user.NewHandler(service)

	return h
}

// UserHandler ...
func (h *handler) UserHandler() user.Handler {
	return h.user
}
