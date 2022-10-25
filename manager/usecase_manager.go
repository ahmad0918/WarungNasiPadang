package manager

import "warung_nasi_padang/usecase"

type UseCaseManager interface {
	MenuUseCase() usecase.MenuUsecase
	TransaksiUseCase() usecase.TransaksiUsecase

}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) MenuUseCase() usecase.MenuUsecase {
	return usecase.NewMenuUseCase(u.repoManager.MenuRepo())
}

func (u *useCaseManager) TransaksiUseCase() usecase.TransaksiUsecase {
	return usecase.NewTransaksiUseCase(u.repoManager.TransaksiRepo())
}


func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}