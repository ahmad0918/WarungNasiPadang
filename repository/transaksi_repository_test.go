package repository

import (
	"database/sql"
	"log"
	"testing"
	"time"
	"warung_nasi_padang/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var dummyTransaksi = []model.Transaksi{
	{
		Id: "C001",
		Menu: "Dummy Menu 1",
		Quantity: 5,
		Date: time.Date(2021,10,11,0,0,0,0,time.Local),
	},
	{
		Id: "C002",
		Menu: "Dummy Menu 2",
		Quantity: 5,
		Date: time.Date(2022,9,16,0,0,0,0,time.Local),
	},
}

type TransaksiRepositoryTestSuite struct {
	suite.Suite
	mockDb *sql.DB
	mockSql sqlmock.Sqlmock
}

func (suite *TransaksiRepositoryTestSuite) SetupTest() {
	mockDb, mockSql, err := sqlmock.New()
	if err != nil {
		log.Fatalln("An error when opening a database connection")
	}
	suite.mockDb = mockDb
	suite.mockSql = mockSql
}

func (suite *TransaksiRepositoryTestSuite) TearDownText() {
	suite.mockDb.Close()
}

func (suite *TransaksiRepositoryTestSuite) TestTransaksiCreate_Success() {
	dummyTransaksi := dummyTransaksi[0]
	suite.mockSql.ExpectExec("insert into transaksi values").WillReturnResult(sqlmock.NewResult(1,1))
	repo := NewTransaksiRepository(suite.mockDb)
	resultRepo := repo.CreateTransaksi(dummyTransaksi)
	assert.Nil(suite.T(), resultRepo)
}

func (suite *TransaksiRepositoryTestSuite) TestTransaksiDetele_success() {
	dummyTransaksi := dummyTransaksi[0]
	suite.mockSql.ExpectExec("delete from transaksi").WillReturnResult(sqlmock.NewResult(1,1))
	repo := NewTransaksiRepository(suite.mockDb)
	resultRepo := repo.DeleteTransaksi(dummyTransaksi.Id)
	assert.Nil(suite.T(), resultRepo)
}

func (suite *TransaksiRepositoryTestSuite) TestTransaksiRetrieveAll_Success() {
	rows := sqlmock.NewRows([]string{"id", "menu", "quantity","date"})
	for _, v := range dummyTransaksi {
		rows.AddRow(v.Id, v.Menu, v.Quantity, v.Date)
	}
	suite.mockSql.ExpectQuery("select \\* from transaksi").WillReturnRows(rows)
	repo := NewTransaksiRepository(suite.mockDb)
	actual, err := repo.ListTransaksi()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(),2, len(actual))
	assert.Equal(suite.T(), "Dummy Menu 1", actual[0].Menu)
}

func (suite *TransaksiRepositoryTestSuite) TestTransaksiUpdate_Success() {
	dummyTransaksi := dummyTransaksi[0]
	suite.mockSql.ExpectExec("update transaksi").WillReturnResult(sqlmock.NewResult(1,1))
	repo := NewTransaksiRepository(suite.mockDb)
	resultRepo := repo.UpdateTransaksi(dummyTransaksi)
	assert.Nil(suite.T(), resultRepo)
}

func TestTransaksiRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TransaksiRepositoryTestSuite))
}