package config

import (
	"errors"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
)

const (
	httpHostEnvName = "HTTP_HOST"
	httpPortEnvName = "HTTP_PORT"

	errMsgHostNotFound = "HTTP host not found"
	errMsgPortNotFound = "HTTP port not found"
)

type HTTPConfig interface {
	Address() string
	ServerConfig(router *chi.Mux) *http.Server
}

type httpConfig struct {
	host string
	port string
}

func NewHTTPConfig() (*httpConfig, error) {
	cfg := &httpConfig{}

	host := os.Getenv(httpHostEnvName)
	if len(host) == 0 {
		return nil, errors.New(errMsgHostNotFound)
	}
	cfg.host = host

	port := os.Getenv(httpPortEnvName)
	if len(port) == 0 {
		return nil, errors.New(errMsgPortNotFound)
	}
	cfg.port = port

	return cfg, nil
}

func (cfg *httpConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}

func (cfg *httpConfig) ServerConfig(router *chi.Mux) *http.Server {
	return &http.Server{
		Addr:         cfg.Address(),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
}
