package services

import (
	"testing"
    "crypto/md5"
    "time"
)

const requestValidatorToken = "super-token"
const anotherRequestValidatorToken = "another-super-token"
var (
    currentTime = time.Now().Unix()
    currentTimeSignature = md5.Sum([]byte(requestValidatorToken + ":" + string(currentTime)))
    anotherTimeSignature = md5.Sum([]byte("signature-value"))
)

func TestCreateNewRequestValidator(t *testing.T) {
    rv := CreateNewRequestValidator(requestValidatorToken)

	if rv == nil {
		t.Error("CreateNewRequestValidator() returns nil")
	}
}

func TestValidate(t *testing.T) {
    rvCurrent := CreateNewRequestValidator(requestValidatorToken)
    rvAnother := CreateNewRequestValidator(anotherRequestValidatorToken)

    if !rvCurrent.Validate(currentTimeSignature, currentTime) {
        t.Error("Signature must be validated properly")
    }

    if rvCurrent.Validate(currentTimeSignature, currentTime + 1) {
        t.Error("Signature must not be validated properly")
    }

    if rvAnother.Validate(anotherTimeSignature, currentTime) {
        t.Error("Signature must not be validated")
    }
}
