package repository

import (
	"warung_nasi_padang/model"
	"warung_nasi_padang/utils"

	"database/sql"
)

type MenuRepository interface {
	Create(newMenu model.Menu) error
	Update(updateMenu model.Menu) error
	Delete(id string)error
	List() ([]model.Menu, error)
}

type menuDbRepository struct {
	db *sql.DB
}

func (m *menuDbRepository) Create(newMenu model.Menu) error {
	_, err := m.db.Exec(utils.INSERT_MENU, newMenu.Id, newMenu.Name, newMenu.Price)
	if err != nil {
		return err
	}
	return nil
}

func (m *menuDbRepository) Update(updateMenu model.Menu) error {
	_, err := m.db.Exec(utils.UPDATE_MENU, updateMenu.Name, updateMenu.Price, updateMenu.Id)
	if err != nil {
		return err
	}
	return nil
}

func (m *menuDbRepository) Delete(id string) error {
	_, err := m.db.Exec(utils.DELETE_MENU,id)
	if err != nil {
		return err
	}
	return nil
}

func (m *menuDbRepository) List() ([]model.Menu, error) {
	var menues []model.Menu
	rows,err := m.db.Query(utils.SELECT_MENU_LIST)
	for rows.Next() {
		menu := model.Menu{}
		err := rows.Scan(&menu.Id, &menu.Name, &menu.Price)
		if err != nil {
			panic(err)
		}
		menues = append(menues, menu)
	}
	err = rows.Err()
	if err != nil{
		return menues, err
	}
		return menues, nil
}

func NewMenuDbRepository(db *sql.DB) MenuRepository {
	repo := new(menuDbRepository)
	repo.db = db
	return repo
}