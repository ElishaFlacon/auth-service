package auth

import (
	"github.com/ElishaFlacon/auth-service/internal/repository"
	def "github.com/ElishaFlacon/auth-service/internal/service"
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
