package Bcrypt

import (
	"BookstoreUtils/errors/rest_errors"

	"golang.org/x/crypto/bcrypt"
)

func Hash(hashString string) (string, *rest_errors.RestErr) {
	hash, err := bcrypt.GenerateFromPassword([]byte(hashString), bcrypt.DefaultCost)
	if err != nil {
		return string(hash), rest_errors.NewInternalServerError(err, "")
	}

	return string(hash), nil
}

func Validate(bcryptString string, validateString string) *rest_errors.RestErr {
	err := bcrypt.CompareHashAndPassword([]byte(bcryptString), []byte(validateString))

	if err != nil {
		return rest_errors.NewUnauthorizedError(err, "unauthorized")
	}

	return nil
}
