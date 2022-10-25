package usecase

import (
	"errors"
	"testing"
	"warung_nasi_padang/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyMenu = []model.Menu{
	{
		Id: "C001",
		Name: "Dummy Name 1",
		Price: 8000,
	},
	{
		Id: "C002",
		Name: "Dummy Name 2",
		Price: 12000,
	},
}

type repoMock struct {
	mock.Mock
}

type MenuUseCaseTestSuite struct {
	suite.Suite
	repoMock *repoMock
}

func (m *repoMock) Create(newMenu model.Menu) error {
	args := m.Called(newMenu)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (m *repoMock) Update(newMenu model.Menu) error {
	args := m.Called(newMenu)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (m *repoMock) Delete(id string) error {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (m *repoMock) List() ([]model.Menu, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Menu), nil
}

func (suite *MenuUseCaseTestSuite) TestMenuCreate_Success() {
	dummy := dummyMenu[0]
	suite.repoMock.On("Create", dummy).Return(nil)
	menuUseCaseTest := NewMenuUseCase(suite.repoMock)
	err := menuUseCaseTest.CreateMenu(dummy)
	assert.Nil(suite.T(), err)
}

func (suite *MenuUseCaseTestSuite) TestMenuCreate_Failed() {
	dummy := dummyMenu[0]
	suite.repoMock.On("Create", dummy).Return(errors.New("failed"))
	menuUseCaseTest := NewMenuUseCase(suite.repoMock)
	err := menuUseCaseTest.CreateMenu(dummy)
	assert.NotNil(suite.T(), err)
}

func (suite *MenuUseCaseTestSuite) TestMenuUpdate_Success() {
	dummy := dummyMenu[0]
	suite.repoMock.On("Update", dummy).Return(nil)
	menuUseCaseTest := NewMenuUseCase(suite.repoMock)
	err := menuUseCaseTest.UpdateMenu(dummy)
	assert.Nil(suite.T(), err)
}

func (suite *MenuUseCaseTestSuite) TestMenuUpdate_Failed() {
	dummy := dummyMenu[0]
	suite.repoMock.On("Update", dummy).Return(errors.New("failed"))
	menuUseCaseTest := NewMenuUseCase(suite.repoMock)
	err := menuUseCaseTest.UpdateMenu(dummy)
	assert.NotNil(suite.T(), err)
}

func (suite *MenuUseCaseTestSuite) TestMenuDelete_Success() {
	dummy := dummyMenu[0]
	suite.repoMock.On("Delete", dummy.Id).Return(nil)
	menuUseCaseTest := NewMenuUseCase(suite.repoMock)
	err := menuUseCaseTest.DeleteMenu(dummy.Id)
	assert.Nil(suite.T(), err)
}

func (suite *MenuUseCaseTestSuite) TestMenuDelete_Failed() {
	dummy := dummyMenu[0]
	suite.repoMock.On("Delete", dummy.Id).Return(errors.New("failed"))
	menuUseCaseTest := NewMenuUseCase(suite.repoMock)
	err := menuUseCaseTest.DeleteMenu(dummy.Id)
	assert.NotNil(suite.T(), err)
}

func (suite *MenuUseCaseTestSuite) TestMenuList_Success() {
	suite.repoMock.On("List").Return(dummyMenu,nil)
	menuUseCaseTest := NewMenuUseCase(suite.repoMock)
	menu, err := menuUseCaseTest.ListMenu()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyMenu, menu)
	assert.NotEmpty(suite.T(), menu)
}

func (suite *MenuUseCaseTestSuite) TestMenuList_Failed() {
	suite.repoMock.On("List").Return(nil,errors.New("failed"))
	menuUseCaseTest := NewMenuUseCase(suite.repoMock)
	menu, err := menuUseCaseTest.ListMenu()
	assert.Error(suite.T(),err)
	assert.NotNil(suite.T(), err)
	assert.Empty(suite.T(), menu)
	assert.Equal(suite.T(), []model.Menu(nil), menu)
}

func (suite *MenuUseCaseTestSuite) SetupTest() {
	suite.repoMock = new(repoMock)
}

func TestMenuUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(MenuUseCaseTestSuite))
}