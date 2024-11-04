package auth

import (
	def "github.com/ElishaFlacon/auth-service/internal/repository"
)

var _ def.AuthRepository = (*repository)(nil)

type repository struct {
	userServicesPath string
}

func NewRepository(userServicesPath string) *repository {
	return &repository{
		userServicesPath: userServicesPath,
	}
}
