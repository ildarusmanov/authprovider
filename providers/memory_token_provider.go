package providers

import (
	"errors"
	"github.com/google/uuid"
	"github.com/ildarusmanov/authprovider/models"
)

// throw the exception when token not found
var tokenNotFound = errors.New("Token not found")

// throw the exception when token duplication detected
var tokenAlreadyExists = errors.New("Token already exists")

// in memory token storege type
type TokensList []string

// in memory token provider type
type MemoryTokenProvider struct {
	tokens     map[string]*models.Token
	userTokens map[string]TokensList
}

// token provider constructor
func CreateNewMemoryTokenProvider() *MemoryTokenProvider {
	p := &MemoryTokenProvider{}
	p.init()

	return p
}

// find a stored token by value
func (p *MemoryTokenProvider) FindByValue(tokenValue string) (*models.Token, error) {
	t, ok := p.tokens[tokenValue]

	if !ok {
		return nil, tokenNotFound
	}

	return t, nil
}

// save token in storage
// return error if the value is alredy exists in storage
func (p *MemoryTokenProvider) AddToken(userId string, scope []string, lifetime int) (*models.Token, error) {
	token := models.CreateNewToken(
		userId,
		p.generateUniqueTokenValue(),
		scope,
		lifetime,
	)

	if t, err := p.FindByValue(token.GetTokenValue()); err == nil {
		return t, tokenAlreadyExists
	}

	p.tokens[token.GetTokenValue()] = token

	p.assignTokenToUser(token.GetTokenUserId(), token.GetTokenValue())

	return token, nil
}

func (p *MemoryTokenProvider) DropToken(tokenValue string) error {
	t, err := p.FindByValue(tokenValue)

	if err != nil {
		return err
	}

	delete(p.tokens, t.GetTokenValue())

	return nil
}

func (p *MemoryTokenProvider) DropByUserId(userId string) error {
	l, err := p.getTokensByUserId(userId)

	if err != nil {
		return err
	}

	for _, t := range l {
		if err = p.DropToken(t); err != nil {
			return err
		}
	}

	return nil
}

// delete all tokens from storage
func (p *MemoryTokenProvider) DropAll() {
	p.init()
}

func (p *MemoryTokenProvider) getTokensByUserId(userId string) (TokensList, error) {
	l, ok := p.userTokens[userId]

	if !ok {
		return nil, tokenNotFound
	}

	return l, nil
}

func (p *MemoryTokenProvider) assignTokenToUser(userId, tokenValue string) {
	_, err := p.getTokensByUserId(userId)

	if err != nil {
		p.userTokens[userId] = TokensList{tokenValue}
	} else {
		p.userTokens[userId] = append(p.userTokens[userId], tokenValue)
	}
}

func (p *MemoryTokenProvider) init() {
	p.tokens = make(map[string]*models.Token)
	p.userTokens = make(map[string]TokensList)
}

func (p *MemoryTokenProvider) generateUniqueTokenValue() string {
	return uuid.New().URN()
}
