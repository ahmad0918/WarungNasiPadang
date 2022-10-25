package repository

import (
	"warung_nasi_padang/model"
	"warung_nasi_padang/utils"

	// "github.com/jmoiron/sqlx"
	"database/sql"
)

type TransaksiRepository interface {
	CreateTransaksi(newTransaksi model.Transaksi) error
	UpdateTransaksi(updateTransaksi model.Transaksi) error
	DeleteTransaksi(id string) error
	ListTransaksi() ([]model.Transaksi, error)
}

type transaksiDbRepository struct {
	db *sql.DB
}

func (t *transaksiDbRepository) CreateTransaksi(newTransaksi model.Transaksi) error {
	_, err := t.db.Exec(utils.INSERT_TRANSAKSI, newTransaksi.Id, newTransaksi.Menu, newTransaksi.Quantity, newTransaksi.Date)
	if err != nil {
		return err
	}
	return nil
}

func (t *transaksiDbRepository) UpdateTransaksi(updateTransaksi model.Transaksi) error {
	_, err := t.db.Exec(utils.UPDATE_TRANSAKSI, updateTransaksi.Menu, updateTransaksi.Quantity, updateTransaksi.Date, updateTransaksi.Id)
	if err != nil {
		return err
	}
	return nil
}

func (t *transaksiDbRepository) DeleteTransaksi(id string) error {
	_, err := t.db.Exec(utils.DELETE_TRANSAKSI,id)
	if err != nil {
		return err
	}
	return nil
}

func (t *transaksiDbRepository) ListTransaksi() ([]model.Transaksi, error) {
	var transaksies []model.Transaksi 
	rows,err := t.db.Query(utils.SELECT_TRANSAKSI_LIST)
	for rows.Next() {
		transaksi := model.Transaksi{}
		err := rows.Scan(&transaksi.Id, &transaksi.Menu, &transaksi.Quantity, &transaksi.Date)
		if err != nil {
			panic(err)
		}
		transaksies = append(transaksies, transaksi)
	}
	err = rows.Err()
	if err != nil{
		return transaksies, err
	}
		return transaksies, nil
}

func NewTransaksiRepository(db *sql.DB) TransaksiRepository {
	repo := new(transaksiDbRepository)
	repo.db = db
	return repo
}