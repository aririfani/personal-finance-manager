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
			router.Action(httppkg.NewRest(http.MethodDelete, "/{id}", handler.AccountHandler().DeleteAccount))
		})

		router.Route("/finance", func(r chi.Router) {
			router := r.(httppkg.Router)
			router.Use(middleware.JWTAuthorization)
			router.Action(httppkg.NewRest(http.MethodPost, "/", handler.FinanceHandler().CreateFinance))
			router.Action(httppkg.NewRest(http.MethodGet, "/", handler.FinanceHandler().GetAllFinance))
			router.Action(httppkg.NewRest(http.MethodPatch, "/{id}", handler.FinanceHandler().UpdateFinance))
			router.Action(httppkg.NewRest(http.MethodGet, "/{id}", handler.FinanceHandler().GetFinanceByID))
			router.Action(httppkg.NewRest(http.MethodDelete, "/{id}", handler.FinanceHandler().DeleteFinanceByID))
			router.Action(httppkg.NewRest(http.MethodGet, "/report/daily", handler.FinanceHandler().GetTotalTransactionDaily))
			router.Action(httppkg.NewRest(http.MethodGet, "/report/monthly", handler.FinanceHandler().GetTotalTransactionMonthly))
		})

		router.Action(httppkg.NewRest(http.MethodPost, "/register", handler.UserHandler().CreateUser))
		router.Action(httppkg.NewRest(http.MethodPost, "/login", handler.UserHandler().Login))

		router.Route("/user", func(r chi.Router) {
			router := r.(httppkg.Router)
			router.Use(middleware.JWTAuthorization)
			router.Action(httppkg.NewRest(http.MethodGet, "/profile", handler.UserHandler().Profile))
		})
	})
	return
}
