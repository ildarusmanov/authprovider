package services

import (
	"github.com/ildarusmanov/authprovider/helpers"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

const (
	requestValidatorToken        = "super-token"
	anotherRequestValidatorToken = "another-super-token"
)

var (
	currentTime              = time.Now().Unix()
	currentTimeSignatureData = requestValidatorToken + ":" + strconv.FormatInt(currentTime, 10)
	currentTimeSignature     = helpers.GetMD5Hash(currentTimeSignatureData)
	anotherTimeSignature     = helpers.GetMD5Hash("signature-value")
)

func TestCreateNewRequestValidator(t *testing.T) {
	rv := CreateNewRequestValidator(requestValidatorToken)

	assert.NotNil(t, rv)
}

func TestValidate(t *testing.T) {
	rvCurrent := CreateNewRequestValidator(requestValidatorToken)
	rvAnother := CreateNewRequestValidator(anotherRequestValidatorToken)

	if !rvCurrent.Validate(currentTimeSignature, currentTime) {
		t.Error("Signature must be validated properly")
	}

	if rvCurrent.Validate(currentTimeSignature, currentTime+1) {
		t.Error("Signature must not be validated properly")
	}

	if rvAnother.Validate(anotherTimeSignature, currentTime) {
		t.Error("Signature must not be validated")
	}
}

func TestCreateSignature(t *testing.T) {
	rv := CreateNewRequestValidator(requestValidatorToken)

	if rv.CreateSignature(currentTime) != currentTimeSignature {
		t.Error("Signatures must be equal")
	}

	if rv.CreateSignature(currentTime+1) == currentTimeSignature {
		t.Error("Signatures must not be equal")
	}
}
