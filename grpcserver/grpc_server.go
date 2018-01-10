package grpcserver

import (
    "golang.org/x/net/context"
    "errors"
)

var invalidReqSignature = errors.New("signatur not valid")

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
        return nil, invalidReqSignature
    }

    return nil, nil
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
