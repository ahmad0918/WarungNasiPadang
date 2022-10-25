package usecase

import (
	"warung_nasi_padang/model"
	"warung_nasi_padang/repository"
)

type MenuUsecase interface {
	CreateMenu(menu model.Menu) error
	UpdateMenu(UpdateMenu model.Menu) error
	DeleteMenu(id string) error
	ListMenu()([]model.Menu, error)
}

type menuUsecase struct {
	repo repository.MenuRepository
}

func (m *menuUsecase) CreateMenu(menu model.Menu) error {
	return m.repo.Create(menu)
}

func (m *menuUsecase) UpdateMenu(UpdateMenu model.Menu) error {
	return m.repo.Update(UpdateMenu)
}

func (m *menuUsecase) DeleteMenu(id string) error {
	return m.repo.Delete(id)
}

func (m *menuUsecase) ListMenu()([]model.Menu, error) {
	return m.repo.List()
}

func NewMenuUseCase(repo repository.MenuRepository) MenuUsecase {
	return &menuUsecase{repo: repo}
}