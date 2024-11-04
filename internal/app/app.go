package app

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

const (
	envPath = ".env"

	msgHTTPServerRunning = "HTTP server is running on %s"
)

type App struct {
	provider   *provider
	httpServer *chi.Mux
}

func NewApp() (*App, error) {
	a := &App{}

	err := a.initDeps()
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	return a.runHTTPServer()
}

func (a *App) initDeps() error {
	inits := []func() error{
		a.initConfig,
		a.initProvider,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f()
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig() error {
	cfg := a.provider.Config(envPath)

	err := cfg.LoadEnv()
	if err != nil {
		return err
	}

	a.provider.HTTPConfig()

	return nil
}

func (a *App) initProvider() error {
	a.provider = newProvider()
	return nil
}

func (a *App) initHTTPServer() error {
	a.httpServer = chi.NewRouter()
	return nil
}

func (a *App) runHTTPServer() error {
	address := a.provider.httpConfig.Address()
	log.Printf(msgHTTPServerRunning, address)

	server := &http.Server{
		Addr:         address,
		Handler:      a.httpServer,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
