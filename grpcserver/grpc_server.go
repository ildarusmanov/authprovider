package grpcserver

import (
	"errors"
	"github.com/ildarusmanov/authprovider/services"
	"golang.org/x/net/context"
)

var invalidReqSignature = errors.New("signature not valid")

// request validator
type requestValidator interface {
	Validate(signature string, timestamp int64) bool
}

// grpc server type
type GrpcServer struct {
	rv requestValidator
	ts *services.TokenService
}

func CreateNewGrpcServer(rv requestValidator, p services.TokenProvider) *GrpcServer {
	return &GrpcServer{rv, services.CreateNewTokenService(p)}
}

func CreateToken(value, userId string, lifetime int32, timestamp int64, scope []string) *Token {
	return &Token{value, userId, lifetime, timestamp, scope}
}

func CreateTokenRequest(signature string, timestamp int64, token *Token) *TokenRequest {
	return &TokenRequest{signature, timestamp, token}
}

func CreateTokenResponse(isOk bool, status string, token *Token) *TokenResponse {
	return &TokenResponse{isOk, status, token}
}

func (s *GrpcServer) AddToken(ctx context.Context, r *TokenRequest) (*TokenResponse, error) {
	if !s.rv.Validate(r.GetSignature(), r.GetTimestamp()) {
		return CreateTokenResponse(false, "fail", nil), invalidReqSignature
	}

	t, err := s.ts.Generate(
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
		return CreateTokenResponse(false, "fail", nil), invalidReqSignature
	}

	return CreateTokenResponse(true, "ok", nil), nil
}

func (s *GrpcServer) DropToken(ctx context.Context, r *TokenRequest) (*TokenResponse, error) {
	if !s.rv.Validate(r.GetSignature(), r.GetTimestamp()) {
		return CreateTokenResponse(false, "fail", nil), invalidReqSignature
	}

	return nil, nil
}

func (s *GrpcServer) ValidateToken(ctx context.Context, r *TokenRequest) (*TokenResponse, error) {
	if !s.rv.Validate(r.GetSignature(), r.GetTimestamp()) {
		return CreateTokenResponse(false, "fail", nil), invalidReqSignature
	}

	return nil, nil
}
