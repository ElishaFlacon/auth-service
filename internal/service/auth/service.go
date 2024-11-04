package auth

import (
	"github.com/ElishaFlacon/shelly/internal/repository"
	def "github.com/ElishaFlacon/shelly/internal/service"
)

var _ def.AuthService = (*service)(nil)

type service struct {
	authRepository repository.AuthRepository
}

func NewService(
	authRepository repository.AuthRepository,
) *service {
	return &service{
		authRepository: authRepository,
	}
}
