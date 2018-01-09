package main

import (
    "testing"
    // "google.golang.org/grpc/test/grpc_testing"
)

func TestCreateNewServer(t *testing.T) {
    s := CreateNewGrpcServer()

    if s == nil {
        t.Error("Empty grpc server")
    }
} 