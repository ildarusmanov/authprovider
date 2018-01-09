package services

import (
    "github.com/ildarusmanov/authprovider/helpers"
)

type RequestValidator struct {
	token string
}

func CreateNewRequestValidator(token string) *RequestValidator {
	return &RequestValidator{token}
}

func (rv *RequestValidator) Validate(signature string, timestamp int64) bool {
    hash := helpers.GetMD5Hash(rv.token + ":" + string(timestamp))

    return hash == signature
}