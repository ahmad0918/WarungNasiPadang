package repository

import (
	"database/sql"
	"log"
	"testing"
	"warung_nasi_padang/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var dummyMenu = []model.Menu{
	{
		Id: "C001",
		Name: "Dummy Name 1",
		Price: 12000,
	},
	{
		Id: "C002",
		Name: "Dummy Name 2",
		Price: 15000,
	},
}

type MenuRepositoryTestSuite struct {
	suite.Suite
	mockDb *sql.DB
	mockSql sqlmock.Sqlmock
}

func (suite *MenuRepositoryTestSuite) SetupTest() {
	mockDb, mockSql, err := sqlmock.New()
	if err != nil {
		log.Fatalln("An error when opening a database connection")
	}
	suite.mockDb = mockDb
	suite.mockSql = mockSql
}

func (suite *MenuRepositoryTestSuite) TearDownText() {
	suite.mockDb.Close()
}

func (suite *MenuRepositoryTestSuite) TestMenuCreate_Success() {
	dummyMenu := dummyMenu[0]
	suite.mockSql.ExpectExec("insert into menu values").WithArgs("C001", "Dummy Name 1", 12000).WillReturnResult(sqlmock.NewResult(1,1))
	repo := NewMenuDbRepository(suite.mockDb)
	resultRepo := repo.Create(dummyMenu)
	assert.Nil(suite.T(), resultRepo)
}

func (suite *MenuRepositoryTestSuite) TestMenuDelete_Success() {
	dummy := dummyMenu[0]
	suite.mockSql.ExpectExec("delete from menu").WillReturnResult(sqlmock.NewResult(1,1))
	repo := NewMenuDbRepository(suite.mockDb)
	resultRepo := repo.Delete(dummy.Id)
	assert.Nil(suite.T(), resultRepo)
}

func (suite *MenuRepositoryTestSuite) TestMenuRetrieveAll_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "price"})
	for _, v := range dummyMenu {
		rows.AddRow(v.Id, v.Name, v.Price)
	}
	suite.mockSql.ExpectQuery("select \\* from menu").WillReturnRows(rows)
	repo := NewMenuDbRepository(suite.mockDb)
	actual, err := repo.List()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(),2, len(actual))
	assert.Equal(suite.T(), "Dummy Name 1", actual[0].Name)
}

func (suite *MenuRepositoryTestSuite) TestMenuUpdate_Success() {
	dummyMenu := dummyMenu[0]
	suite.mockSql.ExpectExec("update menu").WillReturnResult(sqlmock.NewResult(1,1))
	repo := NewMenuDbRepository(suite.mockDb)
	resultRepo := repo.Update(dummyMenu)
	assert.Nil(suite.T(), resultRepo)
}

func TestMenuRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(MenuRepositoryTestSuite))
}