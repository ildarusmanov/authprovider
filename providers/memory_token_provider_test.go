package providers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewMemoryTokenProvider(t *testing.T) {
	p := CreateNewMemoryTokenProvider()

	assert.NotNil(t, p)
}

func TestMemoryTokenProviderAddToken(t *testing.T) {
	var (
		assert         = assert.New(t)
		userId         = "111"
		scopeList      = []string{"scope1", "scope2"}
		otherScopeList = []string{"scope3", "scope4"}
		lifeTime       = 100
	)

	p := CreateNewMemoryTokenProvider()

	token, err := p.AddToken(userId, scopeList, lifeTime)

	assert.Nil(err)

	if assert.NotNil(token) {
		assert.Equal(token.GetTokenUserId(), userId)
		assert.NotEmpty(token.GetTokenValue())
		assert.False(token.InScope(otherScopeList))
		assert.True(token.InScope(scopeList))
	}
}

func TestMemoryTokenProviderFindByValue(t *testing.T) {
	var (
		anotherTokenValue = "token value 123"
		assert            = assert.New(t)
		userId            = "111"
		scopeList         = []string{"scope1", "scope2"}
		lifeTime          = 100
	)

	p := CreateNewMemoryTokenProvider()

	token, err := p.AddToken(userId, scopeList, lifeTime)

	assert.NotNil(token)
	assert.Nil(err)

	storedToken, err := p.FindByValue(token.GetTokenValue())

	assert.Nil(err)

	if assert.NotNil(storedToken) {
		assert.Equal(storedToken.GetTokenValue(), token.GetTokenValue())
	}

	anotherToken, err := p.FindByValue(anotherTokenValue)

	assert.Nil(anotherToken)
	assert.NotNil(err)
}

func TestMemoryTokenProviderDropToken(t *testing.T) {
	var (
		assert    = assert.New(t)
		userId    = "111"
		scopeList = []string{"scope1", "scope2"}
		lifeTime  = 100
	)

	p := CreateNewMemoryTokenProvider()

	token, err := p.AddToken(userId, scopeList, lifeTime)

	assert.Nil(err)
	assert.NotNil(token)

	err = p.DropToken(token.GetTokenValue())

	assert.Nil(err)

	token, err = p.FindByValue(token.GetTokenValue())

	assert.NotNil(err)
	assert.Nil(token)
}

func TestMemoryTokenProviderDropByUserId(t *testing.T) {
	var (
		assert    = assert.New(t)
		userId    = "111"
		scopeList = []string{"scope1", "scope2"}
		lifeTime  = 100
	)

	p := CreateNewMemoryTokenProvider()

	token, err := p.AddToken(userId, scopeList, lifeTime)

	assert.NotNil(token)
	assert.Nil(err)

	p.DropByUserId(token.GetTokenUserId())

	token, err = p.FindByValue(token.GetTokenValue())

	assert.Nil(token)
	assert.NotNil(err)
}

func TestMemoryTokenProviderDropAll(t *testing.T) {
	var (
		assert    = assert.New(t)
		userId    = "111"
		scopeList = []string{"scope1", "scope2"}
		lifeTime  = 100
	)

	p := CreateNewMemoryTokenProvider()

	token, err := p.AddToken(userId, scopeList, lifeTime)

	assert.NotNil(token)
	assert.Nil(err)

	p.DropAll()

	token, err = p.FindByValue(token.GetTokenValue())

	assert.Nil(token)
	assert.NotNil(err)
}
