package services

import (
	"github.com/ildarusmanov/authprovider/providers"
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
		userId           = "111"
		scopeList        = []string{"all"}
		anotherScopeList = []string{"another-scope"}
		lifeTime         = 5
	)

	s := CreateNewTokenService(createTokenProvider())
	s.DropAll()

	token, err := s.Generate(userId, scopeList, lifeTime)

	if err != nil {
		t.Error("Can not generate new token: %s", err)
	}

	if token == nil {
		t.Error("Nil token generated")
	}

	if !token.InScope(scopeList) {
		t.Error("Scope validation failed")
	}

	if token.InScope(anotherScopeList) {
		t.Error("Scope validation failed")
	}

	if token.GetTokenUserId() != userId {
		t.Error("Invalid user id")
	}

	if !token.IsValid() {
		t.Error("Token has expired too fast")
	}

	time.Sleep(time.Duration(lifeTime + 1) * time.Second)

	if token.IsValid() {
		t.Error("Token have to be already expired")
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
