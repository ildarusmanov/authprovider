package models

import (
	"testing"
	"time"
)

func TestCreateNewToken(t *testing.T) {
	var (
		tokenValue     = "token-value-1"
		userId         = "111"
		scope          = []string{"scope1", "scope2"}
		tokenTimestamp = time.Now().Unix()
		lifetime       = 1000
	)

	newToken := CreateNewToken(
		userId,
		tokenValue,
		scope,
		tokenTimestamp,
		lifetime,
	)

	if newToken.GetTokenValue() != tokenValue {
		t.Error("Invalid token value")
	}

	if newToken.GetTokenUserId() != userId {
		t.Error("Invalid user id")
	}

	if newToken.GetTokenScope() != scope {
		t.Error("Invalid token scope")
	}

	if newToken.GetTokenTimestamp() != tokenTimestamp {
		t.Error("Invalid token timestamp")
	}

	if newToken.GetTokenLifetime() != lifetime {
		t.Error("Invalid lifetime")
	}
}

func TestInScope(t *testing.T) {
	var (
		tokenValue     = "token-value-1"
		userId         = "111"
		scope          = []string{"scope1", "scope2"}
		anoherScope    = []string{"scope3", "scope4"}
		tokenTimestamp = time.Now().Unix()
		lifetime       = 1000
	)

	newToken := CreateNewToken(
		userId,
		tokenValue,
		scope,
		tokenTimestamp,
		lifetime,
	)

	if newToken.InScope(anoherScope) {
		t.Error("Wrong scope accepted")
	}

	if !newToken.InScope(scope) {
		t.Error("Valid scope rejected")
	}
}

func TestIsValid(t *testing.T) {
	var (
		tokenValue     = "token-value-1"
		userId         = "111"
		scope          = []string{"scope1", "scope2"}
		anoherScope    = []string{"scope3", "scope4"}
		tokenTimestamp = time.Now().Unix()
		lifetime       = 5
	)

	newToken := CreateNewToken(
		userId,
		tokenValue,
		scope,
		tokenTimestamp,
		lifetime,
	)

	if !newToken.IsValid() {
		t.Error("Token expired too fast")
	}

	time.Sleep((lifetime + 1) * time.Second)

	if newToken.IsValid() {
		t.Error("Token must be expired")
	}
}
