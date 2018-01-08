package services

import (
    "testing"
)

const requestValidatorToken = "super-token"

func TestCreateNewRequestValidator(t *testing.T) {
    rv := CreateNewRequestValidator(requestValidatorToken)

    if rv == nil {
        t.Error("CreateNewRequestValidator() returns nil")
    }
}

func TestValidate(t *testing.T) {
    t.Error("test not implemented yet")   
}
