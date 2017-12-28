package models

type Token struct {
    tokenUserId string
    tokenTimestamp int64
    tokenLifetime int
    tokenScope []string

    GetTokenUserId() string
    GetTokenTimestamp() int64
    GetTokenValue() string
    GetTokenScope() []string
    IsValid() bool
    InScope([]string) bool
}