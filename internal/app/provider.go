package app

import (
	"log"

	"github.com/ElishaFlacon/auth-service/internal/config"
	"github.com/ElishaFlacon/auth-service/internal/controller/auth"
	"github.com/ElishaFlacon/auth-service/internal/repository"
	"github.com/ElishaFlacon/auth-service/internal/service"
	"github.com/go-chi/chi/v5"

	authRepository "github.com/ElishaFlacon/auth-service/internal/repository/auth"
	authService "github.com/ElishaFlacon/auth-service/internal/service/auth"
)

const (
	errMsgGetHTTPConfig       = "Failed to get HTTP config: %s"
	errMsgGetUserServicesPath = "Failed to get user services path: %s"
)

type provider struct {
	config     config.Config
	httpConfig config.HTTPConfig

	authRepository repository.AuthRepository

	authService service.AuthService

	authImpl *auth.Implementation
}

func newProvider() *provider {
	return &provider{}
}

func (p *provider) Config(envPath string) config.Config {
	if p.config == nil {
		cfg := config.NewConfig(envPath)
		p.config = cfg
	}

	return p.config
}

func (p *provider) HTTPConfig() config.HTTPConfig {
	if p.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf(errMsgGetHTTPConfig, err.Error())
		}

		p.httpConfig = cfg
	}

	return p.httpConfig
}

func (p *provider) AuthRepository() repository.AuthRepository {
	if p.authRepository == nil {
		userServicesPath, err := p.config.GetUserServicesPath()
		if err != nil {
			log.Fatalf(errMsgGetUserServicesPath, err.Error())
		}

		p.authRepository = authRepository.NewRepository(userServicesPath)
	}

	return p.authRepository
}

func (p *provider) AuthService() service.AuthService {
	if p.authService == nil {
		repository := p.AuthRepository()
		p.authService = authService.NewService(repository)
	}

	return p.authService
}

func (p *provider) AuthImpl(httpServer *chi.Mux) *auth.Implementation {
	if p.authImpl == nil {
		service := p.AuthService()
		p.authImpl = auth.NewImplementation(
			service,
			httpServer,
		)
	}

	return p.authImpl
}
