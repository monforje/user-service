package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/monforje/user-service/internal/infrastructure/database"
	"github.com/monforje/user-service/internal/repository"
	"github.com/monforje/user-service/internal/service"
	"github.com/monforje/user-service/internal/transport/http"
	"github.com/monforje/user-service/pkg/config"
)

type App struct {
	cfg      *config.Config
	pg       *database.Postgres
	repos    *repository.Repository
	services *service.Service
	handlers *http.Handler
	echo     *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.cfg = config.New()

	a.pg = database.New(a.cfg.Postgres)

	a.repos = repository.New(a.pg.DB)

	a.services = service.New(a.repos)

	a.handlers = http.New(a.services)

	a.echo = echo.New()

	a.echo.HideBanner = true
	a.echo.HidePort = true

	a.echo.Use(middleware.Logger())
	a.echo.Use(middleware.Recover())
	a.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	a.handlers.RegisterRoutes(a.echo)

	return a, nil
}

func (a *App) Run() error {
	log.Println("http: Server start")

	addr := fmt.Sprintf(":%s", a.cfg.App.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := a.echo.Start(addr); err != nil {
			log.Print(err)
		}
	}()

	<-quit
	log.Println("http: Server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.echo.Shutdown(ctx); err != nil {
		log.Print(err)
	}

	a.pg.Stop()

	log.Println("http: Server stopped")
	return nil
}
