package services

import (
	"github.com/ildarusmanov/authprovider/providers"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func createTokenProvider() TokenProvider {
	return providers.CreateNewMemoryTokenProvider()
}

func TestCreateNewTokenService(t *testing.T) {
	s := CreateNewTokenService(createTokenProvider())

	if s == nil {
		t.Error("Service does not created")
	}
}

func TestGenerateToken(t *testing.T) {
	var (
		assert           = assert.New(t)
		userId           = "111"
		scopeList        = []string{"all"}
		anotherScopeList = []string{"another-scope"}
		lifeTime         = 5
	)

	s := CreateNewTokenService(createTokenProvider())
	s.DropAll()

	token, err := s.Generate(userId, scopeList, lifeTime)

	assert.Nil(err)

	if assert.NotNil(token) {
		assert.True(token.InScope(scopeList))
		assert.False(token.InScope(anotherScopeList))
		assert.Equal(token.GetTokenUserId(), userId)
		assert.True(token.IsValid())

		time.Sleep(time.Duration(lifeTime+1) * time.Second)

		assert.False(token.IsValid())
	}
}

func TestValidateToken(t *testing.T) {
	var (
		userId            = "111"
		anotherUserId     = "222"
		scopeList         = []string{"all"}
		lifeTime          = 15
		anotherTokenValue = "another-token-value"
	)

	s := CreateNewTokenService(createTokenProvider())
	s.DropAll()

	token, err := s.Generate(userId, scopeList, lifeTime)

	if userId == anotherUserId {
		t.Error("Users must be different")
	}

	if err != nil {
		t.Error("Can not generate token")
	}

	if !s.Validate(userId, token.GetTokenValue()) {
		t.Error("Generated token is invalid")
	}

	if s.Validate(userId, anotherTokenValue) {
		t.Error("Non-existing token is valid")
	}

	if s.Validate(anotherUserId, token.GetTokenValue()) {
		t.Error("Token can not be valid for other user")
	}
}
