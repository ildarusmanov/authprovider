package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateNewToken(t *testing.T) {
	var (
		assert     = assert.New(t)
		tokenValue = "token-value-1"
		userId     = "111"
		scope      = []string{"scope1", "scope2"}
		lifetime   = 1000
	)

	newToken := CreateNewToken(
		userId,
		tokenValue,
		scope,
		lifetime,
	)

	assert.NotNil(newToken)
	assert.Equal(newToken.GetTokenValue(), tokenValue)
	assert.Equal(newToken.GetTokenUserId(), userId)
	assert.True(newToken.InScope(scope))
	assert.True(newToken.GetTokenTimestamp() <= time.Now().Unix())
	assert.Equal(newToken.GetTokenLifetime(), lifetime)
}

func TestInScope(t *testing.T) {
	var (
		assert      = assert.New(t)
		tokenValue  = "token-value-1"
		userId      = "111"
		scope       = []string{"scope1", "scope2"}
		anoherScope = []string{"scope3", "scope4"}
		lifetime    = 1000
	)

	newToken := CreateNewToken(
		userId,
		tokenValue,
		scope,
		lifetime,
	)

	assert.NotNil(newToken)
	assert.False(newToken.InScope(anoherScope))
	assert.True(newToken.InScope(scope))
}

func TestIsValid(t *testing.T) {
	var (
		assert     = assert.New(t)
		tokenValue = "token-value-1"
		userId     = "111"
		scope      = []string{"scope1", "scope2"}
		lifetime   = 5
	)

	newToken := CreateNewToken(
		userId,
		tokenValue,
		scope,
		lifetime,
	)

	assert.NotNil(newToken)
	assert.True(newToken.IsValid())

	time.Sleep(time.Duration(lifetime+1) * time.Second)

	assert.False(newToken.IsValid())
}
