package server

import (
	"context"
	"github.com/aririfani/personal-finance-manager/internal/app/handler"
	httppkg "github.com/aririfani/personal-finance-manager/internal/pkg/http"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"time"
)

type Server interface {
	Router(handler handler.Handler) (w httppkg.Router)
	GetHTTPServer() *http.Server
	GracefullShutdown(server *http.Server, logger *log.Logger, quit <-chan os.Signal, done chan<- bool)
}

type server struct {
	Addr         string
	Handler      handler.Handler
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func NewServer(addr string, handler handler.Handler, readTimeout time.Duration, writeTimeout time.Duration, idleTimeout time.Duration) Server {
	return &server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}
}

func (s *server) GetHTTPServer() *http.Server {
	return &http.Server{
		Addr:         s.Addr,
		Handler:      s.Router(s.Handler),
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
		IdleTimeout:  s.IdleTimeout,
	}
}

func (s *server) Router(handler handler.Handler) (w httppkg.Router) {
	w = httppkg.NewRouter(chi.NewRouter())
	w.Route("/v1", func(r chi.Router) {
		router := r.(httppkg.Router)
		router.Action(httppkg.NewRest(http.MethodPost, "/user", handler.UserHandler().CreateUser))
		router.Action(httppkg.NewRest(http.MethodPost, "/login", handler.UserHandler().Login))
	})
	return
}

// GracefullShutdown ...
func (s *server) GracefullShutdown(server *http.Server, logger *log.Logger, quit <-chan os.Signal, done chan<- bool) {
	<-quit
	logger.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}
	close(done)
}
