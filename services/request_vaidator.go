package services

import (
    "github.com/ildarusmanov/authprovider/helpers"
    "strconv"
)

type RequestValidator struct {
	token string
}

func CreateNewRequestValidator(token string) *RequestValidator {
	return &RequestValidator{token}
}

func (rv *RequestValidator) Validate(signature string, timestamp int64) bool {
    return rv.CreateSignature(timestamp) == signature
}

func (rv *RequestValidator) CreateSignature(timestamp int64) string {
    return helpers.GetMD5Hash(rv.token + ":" + strconv.FormatInt(timestamp, 10))
}