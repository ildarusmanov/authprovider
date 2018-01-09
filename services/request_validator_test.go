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

	assert := assert.New(t)
	assert.True(rvCurrent.Validate(currentTimeSignature, currentTime))
	assert.False(rvCurrent.Validate(currentTimeSignature, currentTime+1))
	assert.False(rvAnother.Validate(anotherTimeSignature, currentTime))
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
