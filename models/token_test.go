package models

import (
    "testing"
    "time"
)

func TestCreateNewToken(t *testing.T) {
    var (
        tokenValue = "token-value-1"
        userId = "111"
        scope = []string{"scope1", "scope2"}
        tokenTimestamp = time.Now().Unix()
        lifetime = 1000
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
    t.Error("Test does not implemented")
}

func TestIsValid(t *testing.T) {
    t.Error("Test does not implemented")
}
