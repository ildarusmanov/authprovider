package services

type TokenProvider interface {
	FindByValue(tokenValue string) (TokenData, error)
	AddToken(userId string, tokenValue string, scope []string, lifetime int) (TokenData, error)
	DropToken(tokenValue string) error
	DropByUserId(userId string)
	DropAll()
}

type TokenData interface {
	GetTokenUserId() string
	GetTokenTimestamp() int64
	GetTokenValue() string
	GetTokenScope() []fstring
	IsValid() bool
	InScope([]string) bool
}

type TokenService struct {
	provider TokenProvider
}

func CreateNewTokenService(provider TokenProvider) *TokenService {
	return &TokenService{provider}
}

func (s *TokenService) Generate(userId string, scope []string, lifeTime int) (TokenData, error) {
	return s.save(userId, "", scope, lifeTime)
}

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

func (s *TokenService) save(userId, tokenValue string, scope []string, lifetime int) error {
	return s.provider.AddToken(userId, tokenValue, scope, lifetime)
}

func (s *TokenService) find(tokenValue string) (TokenData, error) {
	return s.provider.FindByValue(tokenValue)
}
