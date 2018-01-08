package services

type RequestValdator struct {
    token string
}

func CreateNewRequestValidator(token string) *RequestValdator {
    return&RequestValdator{token}
}
