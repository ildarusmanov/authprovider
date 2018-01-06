package providers

import (
	"errors"
	"github.com/ildarusmanov/authprovider/models"
)

// throw the exception when token not found
var tokenNotFound = errors.New("Token not found")

// throw the exception when token duplication detected
var tokenAlreadyExists = errors.New("Token already exists")

// in memory token storege type
type tokensList []string

// in memory token provider type
type MemoryTokenProvider struct {
	tokens     map[string]*models.Token
	userTokens map[string]tokensList
}

// token provider constructor
func CreateNewMemoryTokenProvider() *MemoryTokenProvider {
	p := &MemoryTokenProvider{}
	p.init()

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

	p.tokens[token.GetTokenValue()] = token

	p.assignTokenToUser(token)

	return token
}

func (p *MemoryTokenProvider) DropToken(tokenValue string) error {
	t, err := p.FindByValue(token.GetTokenValue())

	if err != nil {
		return err
	}

	delete(p.tokens, tokenValue)
}

func (p *MemoryTokenProvider) DropByUserId(userId string) {
	l, err := p.getTokensByUserId(userId)

	if err != nil {
		return
	}

	for u, t := range l {
		if err = p.DropToken(t); err != nil {
			return err
		}
	}
}

// delete all tokens from storage
func (p *MemoryTokenProvider) DropAll() {
	p.init()
}

func (p *MemoryTokenProvider) getTokensByUserId(userId string) (tokensList, error) {
	l, ok := p.userTokens[userId]

	if !ok {
		return nil, error
	}

	return l, nil
}

func (p *MemoryTokenProvider) assignTokenToUser(userId, tokenValue string) {
	l, err := p.getTokensByUserId(userId)

	if err != nil {
		p.userTokens[userId] = make(map[string]string)
	}

	p.userTokens[userId] = tokenValue
}

func (p *MemoryTokenProvider) init() {
	p.tokens = make(map[string]*models.Token)
	p.userTokens = make(map[string]string)
}
