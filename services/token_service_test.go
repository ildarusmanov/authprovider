package services

import (
	"errors"
	"github.com/ildarusmanov/authprovider/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// mokecd token provider type
type tokenProviderMock struct {
	mock.Mock
}

func (m *tokenProviderMock) FindByValue(tokenValue string) (*models.Token, error) {
	args := m.Called(tokenValue)
	return args.Get(0).(*models.Token), args.Error(1)
}

func (m *tokenProviderMock) AddToken(userId string, scope []string, lifetime int) (*models.Token, error) {
	args := m.Called(userId, scope, lifetime)
	return args.Get(0).(*models.Token), args.Error(1)
}

func (m *tokenProviderMock) DropToken(tokenValue string) error {
	args := m.Called(tokenValue)
	return args.Error(0)
}

func (m *tokenProviderMock) DropByUserId(userId string) error {
	args := m.Called(userId)
	return args.Error(0)
}

func (m *tokenProviderMock) DropAll() {
	_ = m.Called()
}

func TestCreateNewTokenService(t *testing.T) {
	p := new(tokenProviderMock)
	s := CreateNewTokenService(p)

	assert.NotNil(t, s)
}

func TestGenerateToken(t *testing.T) {
	var (
		userId     = "111"
		tokenValue = "token value 123"
		scopeList  = []string{"all"}
		lifeTime   = 5
	)

	token := models.CreateNewToken(userId, tokenValue, scopeList, lifeTime)

	p := new(tokenProviderMock)

	p.On("AddToken", userId, scopeList, lifeTime).Return(token, nil)

	gToken, err := CreateNewTokenService(p).Generate(userId, scopeList, lifeTime)

	p.AssertCalled(t, "AddToken", userId, scopeList, lifeTime)

	assert.Nil(t, err)
	assert.Equal(t, gToken, token)
}

func TestValidateToken(t *testing.T) {
	var (
		userId            = "111"
		tokenValue        = "token value 123"
		anotherTokenValue = "another token value 123"
		scopeList         = []string{"all"}
		lifeTime          = 5
	)

	token := models.CreateNewToken(userId, tokenValue, scopeList, lifeTime)

	p := new(tokenProviderMock)

	p.On("FindByValue", tokenValue).Return(token, nil)
	p.On("FindByValue", anotherTokenValue).Return(token, errors.New("error"))

	isValid := CreateNewTokenService(p).Validate(userId, tokenValue)

	assert.True(t, isValid)

	isValid = CreateNewTokenService(p).Validate(userId, anotherTokenValue)

	assert.False(t, isValid)
}

func TestDropToken(t *testing.T) {
	var (
		tokenValue = "token value"
	)
	p := new(tokenProviderMock)

	p.On("DropToken", tokenValue).Return(nil)

	err := CreateNewTokenService(p).DropToken(tokenValue)

	p.AssertCalled(t, "DropToken", tokenValue)

	assert.Nil(t, err)
}

func TestDropByUserId(t *testing.T) {
	var (
		userId = "user id 123"
	)
	p := new(tokenProviderMock)

	p.On("DropByUserId", userId).Return(nil)

	err := CreateNewTokenService(p).DropByUserId(userId)

	p.AssertCalled(t, "DropByUserId", userId)

	assert.Nil(t, err)
}

func TestDropAll(t *testing.T) {
	p := new(tokenProviderMock)

	p.On("DropAll").Return()

	CreateNewTokenService(p).DropAll()

	p.AssertCalled(t, "DropAll")
}
