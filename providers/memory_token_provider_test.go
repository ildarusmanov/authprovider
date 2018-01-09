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
		userId            = "111"
		tokenValue        = "token-value-1"
		anotherTokenValue = "token-value-2"
		scopeList         = []string{"scope1", "scope2"}
		lifeTime          = 100
		newToken          = models.CreateNewToken(userId, tokenValue, scopeList, lifeTime)
	)

	p := CreateNewMemoryTokenProvider()

	token, err := p.AddToken(newToken)

	if err != nil {
		t.Error("Can not add new token")
	}

	storedToken, err := p.FindByValue(tokenValue)

	if err != nil {
		t.Error("Can not find newly added token")
	}

	if storedToken.GetTokenValue() != token.GetTokenValue() {
		t.Error("Tokens are not equal")
	}

	_, err = p.FindByValue(anotherTokenValue)

	if err == nil {
		t.Error("Undefined token can not be found")
	}
}

func TestDropToken(t *testing.T) {
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

	err = p.DropToken(token.GetTokenValue())

	if err != nil {
		t.Error("Can not drop token")
	}

	_, err = p.FindByValue(token.GetTokenValue())

	if err == nil {
		t.Error("Token have not deleted")
	}
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
