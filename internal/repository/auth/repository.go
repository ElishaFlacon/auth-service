package auth

import (
	def "github.com/ElishaFlacon/shelly/internal/repository"
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
