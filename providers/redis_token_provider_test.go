package providers

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func createNewRedisTokenProvider() (*RedisTokenProvider, error) {
    options := map[string]string{}

    return CreateNewRedisTokenProvider(options)
}

func TestCreateRedisTokenProvider(t *testing.T) {
    p, err := createNewRedisTokenProvider(options)

    assert := assert.New(t)
    assert.Nil(err)
    assert.NotNil(p)
}

func TestRedisTokenProviderAddToken(t *testing.T) {
    var (
        assert         = assert.New(t)
        userId         = "111"
        scopeList      = []string{"scope1", "scope2"}
        otherScopeList = []string{"scope3", "scope4"}
        lifeTime       = 100
    )

    p, _ := createNewRedisTokenProvider()()

    token, err := p.AddToken(userId, scopeList, lifeTime)

    assert.Nil(err)

    if assert.NotNil(token) {
        assert.Equal(token.GetTokenUserId(), userId)
        assert.NotEmpty(token.GetTokenValue())
        assert.False(token.InScope(otherScopeList))
        assert.True(token.InScope(scopeList))
    }
}

func TestRedisTokenProviderFindByValue(t *testing.T) {
    var (
        anotherTokenValue = "token value 123"
        assert            = assert.New(t)
        userId            = "111"
        scopeList         = []string{"scope1", "scope2"}
        lifeTime          = 100
    )

    p, _ := createNewRedisTokenProvider()()

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

func TestRedisTokenProviderDropToken(t *testing.T) {
    var (
        assert    = assert.New(t)
        userId    = "111"
        scopeList = []string{"scope1", "scope2"}
        lifeTime  = 100
    )

    p, _ := createNewRedisTokenProvider()()

    token, err := p.AddToken(userId, scopeList, lifeTime)

    assert.Nil(err)
    assert.NotNil(token)

    err = p.DropToken(token.GetTokenValue())

    assert.Nil(err)

    token, err = p.FindByValue(token.GetTokenValue())

    assert.NotNil(err)
    assert.Nil(token)
}

func TestRedisTokenProviderDropByUserId(t *testing.T) {
    var (
        assert    = assert.New(t)
        userId    = "111"
        scopeList = []string{"scope1", "scope2"}
        lifeTime  = 100
    )

    p, _ := createNewRedisTokenProvider()()

    token, err := p.AddToken(userId, scopeList, lifeTime)

    assert.NotNil(token)
    assert.Nil(err)

    p.DropByUserId(token.GetTokenUserId())

    token, err = p.FindByValue(token.GetTokenValue())

    assert.Nil(token)
    assert.NotNil(err)
}

func TestRedisTokenProviderDropAll(t *testing.T) {
    var (
        assert    = assert.New(t)
        userId    = "111"
        scopeList = []string{"scope1", "scope2"}
        lifeTime  = 100
    )

    p, _ := createNewRedisTokenProvider()()

    token, err := p.AddToken(userId, scopeList, lifeTime)

    assert.NotNil(token)
    assert.Nil(err)

    p.DropAll()

    token, err = p.FindByValue(token.GetTokenValue())

    assert.Nil(token)
    assert.NotNil(err)
}
