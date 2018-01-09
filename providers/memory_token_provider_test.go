package providers

import (
	"github.com/ildarusmanov/authprovider/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewMemoryTokenProvider(t *testing.T) {
	p := CreateNewMemoryTokenProvider()

	assert.NotNil(t, p)
}

func TestAddToken(t *testing.T) {
	var (
		assert         = assert.New(t)
		userId         = "111"
		tokenValue     = "token-value-1"
		scopeList      = []string{"scope1", "scope2"}
		otherScopeList = []string{"scope3", "scope4"}
		lifeTime       = 100
		newToken       = models.CreateNewToken(userId, tokenValue, scopeList, lifeTime)
	)

	p := CreateNewMemoryTokenProvider()

	token, err := p.AddToken(newToken)

	assert.Nil(err)

	if assert.NotNil(token) {
		assert.Equal(token.GetTokenUserId(), userId)
		assert.Equal(token.GetTokenValue(), tokenValue)
		assert.False(token.InScope(otherScopeList))
		assert.True(token.InScope(scopeList))
	}
}

func TestFindByValue(t *testing.T) {
	var (
		assert            = assert.New(t)
		userId            = "111"
		tokenValue        = "token-value-1"
		anotherTokenValue = "token-value-2"
		scopeList         = []string{"scope1", "scope2"}
		lifeTime          = 100
		newToken          = models.CreateNewToken(userId, tokenValue, scopeList, lifeTime)
	)

	p := CreateNewMemoryTokenProvider()

	token, err := p.AddToken(newToken)

	assert.NotNil(token)
	assert.Nil(err)

	storedToken, err := p.FindByValue(tokenValue)

	assert.Nil(err)

	if assert.NotNil(storedToken) {
		assert.Equal(storedToken.GetTokenValue(), token.GetTokenValue())
	}

	anotherToken, err := p.FindByValue(anotherTokenValue)

	assert.Nil(anotherToken)
	assert.NotNil(err)
}

func TestDropToken(t *testing.T) {
	var (
		assert     = assert.New(t)
		userId     = "111"
		tokenValue = "token-value-1"
		scopeList  = []string{"scope1", "scope2"}
		lifeTime   = 100
		newToken   = models.CreateNewToken(userId, tokenValue, scopeList, lifeTime)
	)

	p := CreateNewMemoryTokenProvider()

	token, err := p.AddToken(newToken)

	assert.Nil(err)
	assert.NotNil(token)

	err = p.DropToken(token.GetTokenValue())

	assert.Nil(err)

	token, err = p.FindByValue(token.GetTokenValue())

	assert.NotNil(err)
	assert.Nil(token)
}

func TestDropByUserId(t *testing.T) {
	var (
		userId     = "111"
		tokenValue = "token-value-1"
		scopeList  = []string{"scope1", "scope2"}
		lifeTime   = 100
		newToken   = models.CreateNewToken(userId, tokenValue, scopeList, lifeTime)
	)

	p := CreateNewMemoryTokenProvider()

	token, err := p.AddToken(newToken)

	if err != nil {
		t.Error("Can not add new token")
	}

	p.DropByUserId(token.GetTokenUserId())

	_, err = p.FindByValue(token.GetTokenValue())

	if err == nil {
		t.Error("Token have not deleted")
	}
}

func TestDropAll(t *testing.T) {
	var (
		userId     = "111"
		tokenValue = "token-value-1"
		scopeList  = []string{"scope1", "scope2"}
		lifeTime   = 100
		newToken   = models.CreateNewToken(userId, tokenValue, scopeList, lifeTime)
	)

	p := CreateNewMemoryTokenProvider()

	token, err := p.AddToken(newToken)

	if err != nil {
		t.Error("Can not add new token")
	}

	p.DropAll()

	_, err = p.FindByValue(token.GetTokenValue())

	if err == nil {
		t.Error("Token have not deleted")
	}
}
