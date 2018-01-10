package grpcserver

import (
    "github.com/ildarusmanov/authprovider/providers"
    "github.com/ildarusmanov/authprovider/services"
    "golang.org/x/net/context"
    "errors"
)

var invalidReqSignature = errors.New("signature not valid")

type requestValidator interface {
    Validate(signature string, timestamp int64) bool
}

type GrpcServer struct{
    rv requestValidator
}

func CreateNewGrpcServer(rv requestValidator) *GrpcServer {
	return &GrpcServer{rv}
}

func CreateToken(value, userId string, lifetime int32, timestamp int64, scope []string) *Token {
    return &Token{value, userId, lifetime, timestamp, scope}
}

func CreateTokenRequest(signature string, timestamp int64, token *Token)  *TokenRequest {
    return &TokenRequest{signature, timestamp, token}
}

func CreateTokenResponse(isOk bool, status string, token *Token)  *TokenResponse{
    return &TokenResponse{isOk, status, token}
}

func (s *GrpcServer) AddToken(ctx context.Context, r *TokenRequest) (*TokenResponse, error) {
    if !s.rv.Validate(r.GetSignature(), r.GetTimestamp()) {
        return CreateTokenResponse(false, "fail", nil), invalidReqSignature
    }

    t, err := createNewTokenService().Generate(
        r.GetToken().GetUserId(),
        r.GetToken().GetScope(),
        int(r.GetToken().GetLifetime()),
    )

    if err != nil {
        return CreateTokenResponse(false, "fail", nil), err
    }

    token := CreateToken(
        t.GetTokenValue(),
        t.GetTokenUserId(),
        int32(t.GetTokenLifetime()),
        t.GetTokenTimestamp(),
        t.GetTokenScope(),
    )

    return CreateTokenResponse(true, "ok", token), nil
}

func (s *GrpcServer) FindToken(ctx context.Context, r *TokenRequest) (*TokenResponse, error) {
    if !s.rv.Validate(r.GetSignature(), r.GetTimestamp()) {
        return nil, invalidReqSignature
    }

    return nil, nil
}

func (s *GrpcServer) DropToken(ctx context.Context, r *TokenRequest) (*TokenResponse, error) {
    if !s.rv.Validate(r.GetSignature(), r.GetTimestamp()) {
        return nil, invalidReqSignature
    }

    return nil, nil
}

func (s *GrpcServer) ValidateToken(ctx context.Context, r *TokenRequest) (*TokenResponse, error) {
    if !s.rv.Validate(r.GetSignature(), r.GetTimestamp()) {
        return nil, invalidReqSignature
    }

    return nil, nil
}

func createNewTokenService() *services.TokenService {
    p := providers.CreateNewMemoryTokenProvider()

    return services.CreateNewTokenService(p)
}
