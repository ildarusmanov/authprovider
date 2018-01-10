package grpcserver

import (
	"github.com/stretchr/testify/assert"
	"testing"
	// "google.golang.org/grpc/test/grpc_testing"
)

func TestCreateNewServer(t *testing.T) {
	s := CreateNewGrpcServer()

	assert.NotNil(t, s)
    assert.Implements(t, (*TokenStorageServer)(nil), s)
}
