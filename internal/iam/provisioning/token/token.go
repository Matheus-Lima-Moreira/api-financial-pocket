package token

import (
	"crypto/rand"
	"encoding/base64"

	shared_errors "github.com/Matheus-Lima-Moreira/financial-pocket/internal/shared/errors"
)

func GenerateToken() (string, *shared_errors.AppError) {
	b := make([]byte, 32)

	_, err := rand.Read(b)
	if err != nil {
		return "", shared_errors.NewBadRequest(err.Error())
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}
