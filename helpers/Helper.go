package helpers

import (
	"crypto/rand"
	"golang.org/x/crypto/bcrypt"
	"simkes-go/models/response"
)

func GetResponse(status int, data interface{}, err error) response.Response {
	var (
		response response.Response
	)

	switch status {
	case 200:
		response.Message = "Success"
	default:
		response.Message = err.Error()
	}

	response.Status = status
	response.Data = data

	return response
}

func GenerateSalt() []byte {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		panic(err)
	}
	return salt
}

func HashPassword(password string, salt []byte) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
