package config

import (
	"errors"
	"net"
	"os"
)

const (
	httpHostEnvName = "HTTP_HOST"
	httpPortEnvName = "HTTP_PORT"

	errMsgHostNotFound = "HTTP host not found"
	errMsgPortNotFound = "HTTP port not found"
)

type HTTPConfig interface {
	Address() string
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
