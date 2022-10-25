package manager

import (
	"warung_nasi_padang/repository"
	"warung_nasi_padang/utils/authenticator"
)

type RepositoryManager interface {
	MenuRepo() repository.MenuRepository
	TransaksiRepo() repository.TransaksiRepository
}

type repositoryManager struct {
	infra InfraManager
	tokenService authenticator.AccessToken
}

func (r *repositoryManager) MenuRepo() repository.MenuRepository  {
	return repository.NewMenuDbRepository(r.infra.SqlDB())
}

func (r *repositoryManager) TransaksiRepo() repository.TransaksiRepository  {
	return repository.NewTransaksiRepository(r.infra.SqlDB())
}

func NewRepositoryManager(manager InfraManager) RepositoryManager {
	return &repositoryManager{infra: manager}
}