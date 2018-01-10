package services

import (
	"github.com/ildarusmanov/authprovider/models"
)

// interface for token provider
type TokenProvider interface {
	FindByValue(tokenValue string) (*models.Token, error)
	AddToken(userId string, scope []string, lifetime int) (*models.Token, error)
	DropToken(tokenValue string) error
	DropByUserId(userId string) error
	DropAll()
}

// token service type
type TokenService struct {
	provider TokenProvider
}

// create new token service
func CreateNewTokenService(provider TokenProvider) *TokenService {
	return &TokenService{provider}
}

// generate new token value
func (s *TokenService) Generate(userId string, scope []string, lifeTime int) (*models.Token, error) {
	return s.provider.AddToken(userId, scope, lifeTime)
}

// validate token value for given user
func (s *TokenService) Validate(userId string, tokenValue string) bool {
	token, err := s.find(tokenValue)

	if err != nil {
		return false
	}

	if token.GetTokenUserId() != userId {
		return false
	}

	return token.IsValid()
}

func (s *TokenService) DropToken(tokenValue string) error {
	return s.provider.DropToken(tokenValue)
}

func (s *TokenService) DropByUserId(userId string) error {
	return s.provider.DropByUserId(userId)
}

func (s *TokenService) DropAll() {
	s.provider.DropAll()
}

// find token by value
func (s *TokenService) find(tokenValue string) (*models.Token, error) {
	return s.provider.FindByValue(tokenValue)
}
