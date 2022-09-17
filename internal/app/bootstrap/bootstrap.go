package bootstrap

import (
	"github.com/aririfani/personal-finance-manager/config"
	"github.com/aririfani/personal-finance-manager/internal/app"
	"github.com/aririfani/personal-finance-manager/internal/app/appcontext"
	"github.com/aririfani/personal-finance-manager/internal/app/handler"
	"github.com/aririfani/personal-finance-manager/internal/app/server"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Run() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	cfg := config.NewConfig()
	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	appContext := appcontext.NewAppContext()

	db, err := appContext.GetDB()
	if err != nil {
		log.Fatal(err)
	}

	repo := app.WiringRepository(db)
	srv := app.WiringService(repo)

	s := server.NewServer(
		net.JoinHostPort(cfg.GetString("server.host"), cfg.GetString("server.port")),
		handler.NewHandler(srv),
		time.Duration(cfg.GetInt("server.read_timeout"))*time.Second,
		time.Duration(cfg.GetInt("server.write_timeout"))*time.Second,
		time.Duration(cfg.GetInt("server.idle_timeout"))*time.Second,
	)

	httpServer := s.GetHTTPServer()
	go s.GracefullShutdown(httpServer, logger, quit, done)

	logger.Println("=> http server started on", net.JoinHostPort(cfg.GetString("server.host"), cfg.GetString("server.port")))
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", cfg.GetString("server.port"), err)
	}

	<-done

	logger.Println("Server stopped")
}
