package app

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	envPath = ".env"

	msgHTTPServerRunning = "HTTP server is running on %s"
)

type App struct {
	provider   *provider
	httpServer *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	return a.runHTTPServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initProvider,
		a.initConfig,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initProvider(_ context.Context) error {
	a.provider = newProvider()
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	cfg := a.provider.Config(envPath)

	err := cfg.LoadEnv()
	if err != nil {
		return err
	}

	a.provider.HTTPConfig()

	return nil
}

func (a *App) initMiddlewares(router *chi.Mux) {
	router.Use(middleware.Logger)
}

func (a *App) initRoutes(router *chi.Mux) {
	a.provider.AuthImpl(router).RegisterRoutes()
}

func (a *App) initHTTPServer(_ context.Context) error {
	router := chi.NewRouter()

	a.initMiddlewares(router)
	a.initRoutes(router)

	a.httpServer = a.provider.httpConfig.ServerConfig(router)

	// TODO: graceful shutdown

	return nil
}

func (a *App) runHTTPServer() error {
	address := a.provider.httpConfig.Address()
	log.Printf(msgHTTPServerRunning, address)

	err := a.httpServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
