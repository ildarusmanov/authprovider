package grpcserver

import (
    "github.com/ildarusmanov/authprovider/services"
	"github.com/stretchr/testify/assert"
    "golang.org/x/net/context"
	"testing"
    "time"
	// "google.golang.org/grpc/test/grpc_testing"
)

const rvToken = "request validator token"

func TestCreateNewServer(t *testing.T) {
    rv := services.CreateNewRequestValidator(rvToken)
	s := CreateNewGrpcServer(rv)

	assert.NotNil(t, s)
    assert.Implements(t, (*TokenStorageServer)(nil), s)
}

func TestAddToken(t *testing.T) {
    assert := assert.New(t)

    token := CreateToken(
        "",
        "123",
        100,
        time.Now().Unix(),
        []string{"all"},
    )

    rv := services.CreateNewRequestValidator(rvToken)
    s := CreateNewGrpcServer(rv)
    timestamp := time.Now().Unix()
    signature := rv.CreateSignature(timestamp)

    validReq := CreateTokenRequest(signature, timestamp, token)
    invalidReq := CreateTokenRequest("some-sig", timestamp, nil)

    invalidResp, err := s.AddToken(context.Background(), invalidReq)
    assert.NotNil(err)
    assert.False(invalidResp.GetIsOk())

    validResp, err := s.AddToken(context.Background(), validReq)
    assert.Nil(err)
    assert.True(validResp.GetIsOk())
    assert.Equal(token.GetUserId(), validResp.GetToken().GetUserId())
}
