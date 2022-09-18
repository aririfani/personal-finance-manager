package server

import (
	"github.com/aririfani/personal-finance-manager/internal/app/handler"
	"github.com/aririfani/personal-finance-manager/internal/app/middleware"
	httppkg "github.com/aririfani/personal-finance-manager/internal/pkg/http"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (s *server) Router(handler handler.Handler) (w httppkg.Router) {
	w = httppkg.NewRouter(chi.NewRouter())
	w.Route("/v1", func(r chi.Router) {
		router := r.(httppkg.Router)
		router.Route("/account", func(r chi.Router) {
			router := r.(httppkg.Router)
			router.Use(middleware.JWTAuthorization)
			router.Action(httppkg.NewRest(http.MethodPost, "/", handler.AccountHandler().CreateAccount))
			router.Action(httppkg.NewRest(http.MethodPatch, "/{id}", handler.AccountHandler().UpdateAccount))
			router.Action(httppkg.NewRest(http.MethodGet, "/", handler.AccountHandler().GetAllAccount))
			router.Action(httppkg.NewRest(http.MethodGet, "/{id}", handler.AccountHandler().GetByAccountByID))
		})

		router.Route("/finance", func(r chi.Router) {
			router := r.(httppkg.Router)
			router.Use(middleware.JWTAuthorization)
			router.Action(httppkg.NewRest(http.MethodPost, "/", handler.FinanceHandler().CreateFinance))
			router.Action(httppkg.NewRest(http.MethodGet, "/", handler.FinanceHandler().GetAllFinance))
			router.Action(httppkg.NewRest(http.MethodPatch, "/{id}", handler.FinanceHandler().UpdateFinance))
		})

		router.Action(httppkg.NewRest(http.MethodPost, "/user", handler.UserHandler().CreateUser))
		router.Action(httppkg.NewRest(http.MethodPost, "/login", handler.UserHandler().Login))
	})
	return
}
