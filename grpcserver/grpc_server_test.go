package grpcserver

import (
	"github.com/ildarusmanov/authprovider/providers"
	"github.com/ildarusmanov/authprovider/services"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"testing"
	"time"
	// "google.golang.org/grpc/test/grpc_testing"
)

const rvToken = "request validator token"

func TestCreateToken(t *testing.T) {
	var (
		tValue     = "ttt"
		tUserId    = "123"
		tLifetime  = int32(100)
		tTimestamp = time.Now().Unix()
		tScope     = []string{"all"}
	)

	token := CreateToken(
		tValue,
		tUserId,
		tLifetime,
		tTimestamp,
		tScope,
	)

	assert := assert.New(t)

	if assert.NotNil(token) {
		assert.Equal(tValue, token.GetValue())
		assert.Equal(tUserId, token.GetUserId())
		assert.Equal(tLifetime, token.GetLifetime())
		assert.Equal(tScope, token.GetScope())
	}
}

func TestCreateTokenResponse(t *testing.T) {
	var (
		rIsOk   = true
		rStatus = "status-text"
		rToken  = CreateToken(
			"ttt",
			"123",
			100,
			time.Now().Unix(),
			[]string{"all"},
		)
	)

	resp := CreateTokenResponse(rIsOk, rStatus, rToken)

	assert := assert.New(t)
	assert.Equal(resp.GetIsOk(), rIsOk)
	assert.Equal(resp.GetStatus(), rStatus)

	if assert.NotNil(resp.GetToken()) {
		assert.Equal(resp.GetToken().GetValue(), rToken.GetValue())
		assert.Equal(resp.GetToken().GetUserId(), rToken.GetUserId())
		assert.Equal(resp.GetToken().GetLifetime(), rToken.GetLifetime())
		assert.Equal(resp.GetToken().GetTimestamp(), rToken.GetTimestamp())
		assert.Equal(resp.GetToken().GetScope(), rToken.GetScope())
	}

}

func TestCreateTokenRequest(t *testing.T) {
	var (
		rSignature = "signature"
		rTimestamp = time.Now().Unix()
		rToken     = CreateToken(
			"ttt",
			"123",
			100,
			time.Now().Unix(),
			[]string{"all"},
		)
	)

	req := CreateTokenRequest(rSignature, rTimestamp, rToken)

	assert := assert.New(t)
	assert.Equal(req.GetSignature(), rSignature)
	assert.Equal(req.GetTimestamp(), rTimestamp)

	// validate token
	if assert.NotNil(req.GetToken()) {
		assert.Equal(req.GetToken().GetValue(), rToken.GetValue())
		assert.Equal(req.GetToken().GetUserId(), rToken.GetUserId())
		assert.Equal(req.GetToken().GetLifetime(), rToken.GetLifetime())
		assert.Equal(req.GetToken().GetTimestamp(), rToken.GetTimestamp())
		assert.Equal(req.GetToken().GetScope(), rToken.GetScope())
	}
}

func TestCreateNewServer(t *testing.T) {
	rv := services.CreateNewRequestValidator(rvToken)
	p := providers.CreateNewMemoryTokenProvider()
	s := CreateNewGrpcServer(rv, p)

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
	p := providers.CreateNewMemoryTokenProvider()
	s := CreateNewGrpcServer(rv, p)
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

func TestDropToken(t *testing.T) {
	assert := assert.New(t)

	var (
		userId1    = "123"
		userId2    = "134"
		tokenValue = "1"
		lifetime   = 100
		scope      = []string{"all"}
	)

	rv := services.CreateNewRequestValidator(rvToken)
	p := providers.CreateNewMemoryTokenProvider()

	token, err := p.AddToken(userId1, scope, lifetime)

	assert.Nil(err)

	token1 := CreateToken(
		token.GetTokenValue(),
		token.GetTokenUserId(),
		int32(token.GetTokenLifetime()),
		token.GetTokenTimestamp(),
		token.GetTokenScope(),
	)

	token2 := CreateToken(
		tokenValue,
		userId2,
		int32(lifetime),
		time.Now().Unix(),
		scope,
	)

	s := CreateNewGrpcServer(rv, p)
	timestamp := time.Now().Unix()
	signature := rv.CreateSignature(timestamp)

	validReq := CreateTokenRequest(signature, timestamp, token1)
	invalidReq := CreateTokenRequest(signature, timestamp, token2)

	invalidResp, err := s.DropToken(context.Background(), invalidReq)
	assert.NotNil(err)
	assert.False(invalidResp.GetIsOk())

	validResp, err := s.DropToken(context.Background(), validReq)
	assert.Nil(err)
	assert.True(validResp.GetIsOk())
}

func TestValidateToken(t *testing.T) {
	assert := assert.New(t)

	var (
		userId1    = "123"
		userId2    = "123"
		tokenValue = "1"
		lifetime   = 100
		scope      = []string{"all"}
	)

	rv := services.CreateNewRequestValidator(rvToken)
	p := providers.CreateNewMemoryTokenProvider()

	token, err := p.AddToken(userId1, scope, lifetime)

	assert.Nil(err)

	token1 := CreateToken(
		token.GetTokenValue(),
		token.GetTokenUserId(),
		int32(token.GetTokenLifetime()),
		token.GetTokenTimestamp(),
		token.GetTokenScope(),
	)

	token2 := CreateToken(
		tokenValue,
		userId2,
		int32(lifetime),
		time.Now().Unix(),
		scope,
	)

	s := CreateNewGrpcServer(rv, p)
	timestamp := time.Now().Unix()
	signature := rv.CreateSignature(timestamp)

	validReq := CreateTokenRequest(signature, timestamp, token1)
	invalidReq := CreateTokenRequest(signature, timestamp, token2)

	invalidResp, err := s.ValidateToken(context.Background(), invalidReq)
	assert.NotNil(err)
	assert.False(invalidResp.GetIsOk())

	validResp, err := s.ValidateToken(context.Background(), validReq)
	assert.Nil(err)
	assert.True(validResp.GetIsOk())
}
