package models

import (
	"time"
)

type Token struct {
	tokenUserId    string
	tokenValue     string
	tokenScope     []string
	tokenTimestamp int64
	tokenLifetime  int
	tokenScopeMap  map[string]struct{}
}

func CreateNewToken(tokenUserId, tokenValue string, tokenScope []string, tokenLifetime int) *Token {
	newToken := &Token{
		tokenUserId,
		tokenValue,
		tokenScope,
		time.Now().Unix(),
		tokenLifetime,
		nil,
	}

	newToken.initScope()

	return newToken
}

func (t *Token) initScope() {
	t.tokenScopeMap = make(map[string]struct{}, len(t.tokenScope))

	for _, s := range t.tokenScope {
		t.tokenScopeMap[s] = struct{}{}
	}
}

func (t *Token) GetTokenUserId() string {
	return t.tokenUserId
}

func (t *Token) GetTokenValue() string {
	return t.tokenValue
}

func (t *Token) GetTokenTimestamp() int64 {
	return t.tokenTimestamp
}

func (t *Token) GetTokenLifetime() int {
	return t.tokenLifetime
}

func (t *Token) GetTokenScope() []string {
	return t.tokenScope
}

func (t *Token) InScope(needle []string) bool {
	allInScope := true

	for _, s := range needle {
		if _, ok := t.tokenScopeMap[s]; !ok {
			allInScope = false
		}
	}

	return allInScope
}

func (t *Token) IsValid() bool {
	if t.tokenLifetime == 0 {
		return true
	}

	return time.Now().Unix() < t.tokenTimestamp+int64(t.tokenLifetime)
}
