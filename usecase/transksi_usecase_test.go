package usecase

import (
	"errors"
	"testing"
	"time"
	"warung_nasi_padang/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

type repoTransaksiMock struct{
	mock.Mock
}

type TransaksiUseCaseTestSuite struct {
	suite.Suite
	repoMock *repoTransaksiMock
}

func (t *repoTransaksiMock) CreateTransaksi(newTransaksi model.Transaksi) error {
	args := t.Called(newTransaksi)
	if args.Get(0) != nil {
		return args.Error(0)
	} 
	return nil
}

func (t *repoTransaksiMock) UpdateTransaksi(newTransaksi model.Transaksi) error {
	args := t.Called(newTransaksi)
	if args.Get(0) != nil {
		return args.Error(0)
	} 
	return nil
}

func (t *repoTransaksiMock) DeleteTransaksi(id string) error {
	args := t.Called(id)
	if args.Get(0) != nil {
		return args.Error(0)
	} 
	return nil
}

func (t *repoTransaksiMock) ListTransaksi() ([]model.Transaksi, error) {
	args := t.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Transaksi), nil
}

func (suite *TransaksiUseCaseTestSuite) TestTransaksiCreate_Success() {
	dummy := dummyTransaksi[0]
	suite.repoMock.On("CreateTransaksi", dummy).Return(nil)
	transaksiUseCaseTest := NewTransaksiUseCase(suite.repoMock)
	err := transaksiUseCaseTest.CreateTransaksi(dummy)
	assert.Nil(suite.T(),err)
}

func (suite *TransaksiUseCaseTestSuite) TestTransaksiCreate_Failed() {
	dummy := dummyTransaksi[0]
	suite.repoMock.On("CreateTransaksi", dummy).Return(errors.New("failed"))
	transaksiUsecaseTest := NewTransaksiUseCase(suite.repoMock)
	err := transaksiUsecaseTest.CreateTransaksi(dummy)
	assert.NotNil(suite.T(), err)
}

func (suite *TransaksiUseCaseTestSuite) TestTrasaksiUpdate_Success() {
	dummy := dummyTransaksi[0]
	suite.repoMock.On("UpdateTransaksi", dummy).Return(nil)
	transaksiUseCaseTest := NewTransaksiUseCase(suite.repoMock)
	err := transaksiUseCaseTest.UpdateTransaksi(dummy)
	assert.Nil(suite.T(),err)
}


func (suite *TransaksiUseCaseTestSuite) TestTransaksiUpdate_Failed() {
	dummy := dummyTransaksi[0]
	suite.repoMock.On("UpdateTransaksi", dummy).Return(errors.New("failed"))
	transaksiUseCaseTest := NewTransaksiUseCase(suite.repoMock)
	err := transaksiUseCaseTest.UpdateTransaksi(dummy)
	assert.NotNil(suite.T(),err)
}

func (suite *TransaksiUseCaseTestSuite) TestTransaksiDelete_Success() {
	dummy := dummyTransaksi[0]
	suite.repoMock.On("DeleteTransaksi", dummy.Id).Return(nil)
	transaksiUseCaseTest := NewTransaksiUseCase(suite.repoMock)
	err := transaksiUseCaseTest.DeleteTransaksi(dummy.Id)
	assert.Nil(suite.T(),err)
}

func (suite *TransaksiUseCaseTestSuite) TestTransaksiDelete_Failed() {
	dummy := dummyTransaksi[0]
	suite.repoMock.On("DeleteTransaksi", dummy.Id).Return(errors.New("failed"))
	transaksiUseCaseTest := NewTransaksiUseCase(suite.repoMock)
	err := transaksiUseCaseTest.DeleteTransaksi(dummy.Id)
	assert.NotNil(suite.T(), err)
}

func (suite *TransaksiUseCaseTestSuite) TestTransaksiList_Failed() {
	suite.repoMock.On("ListTransaksi").Return(nil,errors.New("failed"))
	transaksiUseCaseTest := NewTransaksiUseCase(suite.repoMock)
	transaksi,err := transaksiUseCaseTest.ListTransaksi()
	assert.Error(suite.T(),err)
	assert.NotNil(suite.T(),err)
	assert.Empty(suite.T(), transaksi)
	assert.Equal(suite.T(), []model.Transaksi(nil), transaksi)
}

func (suite *TransaksiUseCaseTestSuite) SetupTest() {
	suite.repoMock = new(repoTransaksiMock)
}

func TestTransaksiUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(TransaksiUseCaseTestSuite))
}