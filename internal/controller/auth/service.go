package auth

import (
	"github.com/ElishaFlacon/auth-service/internal/service"
	"github.com/go-chi/chi/v5"
)

type Implementation struct {
	authService service.AuthService
	httpServer  *chi.Mux
}

func NewImplementation(
	authService service.AuthService,
	httpServer *chi.Mux,
) *Implementation {
	return &Implementation{
		authService: authService,
		httpServer:  httpServer,
	}
}

func (i *Implementation) RegisterRoutes() {
	i.httpServer.Post("/register", i.Register)
	i.httpServer.Post("/login", i.Login)
	i.httpServer.Post("/logout", i.Logout)
	i.httpServer.Get("/check-auth", i.CheckAuth)
}
