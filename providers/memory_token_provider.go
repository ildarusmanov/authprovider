package providers

import (
    "github.com/ildarusmanov/authprovider/models"
    "errors"
)

var tokenNotFound = errors.New("Token not found")
var tokenAlreadyExists = errors.New("Token already exists")

type MemoryTokenProvider struct {
	tokens map[string]*models.Token
}

func CreateNewMemoryTokenProvider() *MemoryTokenProvider {
    p := &MemoryTokenProvider{}

    return p
}

func (p *MemoryTokenProvider) FindByValue(tokenValue string) (*models.Token, error) {
    t, ok := p.tokens[tokenValue]

    if !ok {
        return nil, tokenNotFound
    }

    return t, nil
}

func (p *MemoryTokenProvider) AddToken(token *models.Token) (*models.Token, error) {
    t, err := p.FindByValue(token.GetTokenValue())

    if err != nil {
        return nil, err
    }

    return 
}

func (p *MemoryTokenProvider) DropToken(tokenValue string) error {
    
}

func (p *MemoryTokenProvider) DropByUserId(userId string) {
    
}

func (p *MemoryTokenProvider) DropAll()) {
    
}
