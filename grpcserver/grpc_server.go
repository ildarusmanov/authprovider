package grpcserver

import (
    "golang.org/x/net/context"
)

type GrpcServer struct{}

func CreateNewGrpcServer() *GrpcServer {
	return &GrpcServer{}
}

func (s *GrpcServer) AddToken(context.Context, *TokenRequest) (*TokenResponse, error) {
    return nil, nil
}

func (s *GrpcServer) FindToken(context.Context, *TokenRequest) (*TokenResponse, error) {
    return nil, nil
}

func (s *GrpcServer) DropToken(context.Context, *TokenRequest) (*TokenResponse, error) {
    return nil, nil
}

func (s *GrpcServer) ValidateToken(context.Context, *TokenRequest) (*TokenResponse, error) {
    return nil, nil
}
