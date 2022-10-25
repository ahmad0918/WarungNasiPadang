package usecase

import (
	"warung_nasi_padang/model"
	"warung_nasi_padang/repository"
)

type TransaksiUsecase interface {
	CreateTransaksi(transaksi model.Transaksi) error
	UpdateTransaksi(updateTransaksi model.Transaksi) error
	DeleteTransaksi(id string) error
	ListTransaksi()([]model.Transaksi, error)
}

type transaksiUsecase struct {
	repo repository.TransaksiRepository
}

func (t *transaksiUsecase) CreateTransaksi(transaksi model.Transaksi) error {
	return t.repo.CreateTransaksi(transaksi)
}

func (t *transaksiUsecase) UpdateTransaksi(updateTransaksi model.Transaksi) error {
	return t.repo.UpdateTransaksi(updateTransaksi)
}

func (t *transaksiUsecase) DeleteTransaksi(id string) error {
	return t.repo.DeleteTransaksi(id)
}

func (t *transaksiUsecase) ListTransaksi()([]model.Transaksi, error) {
	return t.repo.ListTransaksi()
}

func NewTransaksiUseCase(repo repository.TransaksiRepository) TransaksiUsecase {
	return &transaksiUsecase{repo: repo}
}