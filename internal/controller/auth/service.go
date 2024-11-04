package auth

import "github.com/ElishaFlacon/shelly/internal/service"

type Implementation struct {
	authService service.AuthService
}

func NewImplementation(authService service.AuthService) *Implementation {
	return &Implementation{
		authService: authService,
	}
}
