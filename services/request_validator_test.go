package services

import (
	"testing"
    "time"
    "strconv"
    "github.com/ildarusmanov/authprovider/helpers"
)

const (
    requestValidatorToken = "super-token"
    anotherRequestValidatorToken = "another-super-token"
)

var (
    currentTime = time.Now().Unix()
    currentTimeSignatureData = requestValidatorToken + ":" + strconv.FormatInt(currentTime, 10)
    currentTimeSignature = helpers.GetMD5Hash(currentTimeSignatureData)
    anotherTimeSignature = helpers.GetMD5Hash("signature-value")
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
