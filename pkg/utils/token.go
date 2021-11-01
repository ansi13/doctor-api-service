package utils

import (
	"log"
	"net/http"
	"time"

	"github.com/pascaldekloe/jwt"
)

func GenerateToken(username string, secret []byte) (string, error) {
	var claims jwt.Claims

	payload := map[string]interface{}{
		"username": username,
	}

	claims.Subject = "User validation"
	claims.Issued = jwt.NewNumericTime(time.Now().Round(time.Second))
	claims.Expires = jwt.NewNumericTime(time.Now().Round(time.Second).AddDate(0, 0, 2))
	claims.Set = payload

	token, err := claims.HMACSign("HS256", secret)

	return string(token), err
}

func ValidateToken(r *http.Request, secret []byte) (string, error) {
	claims, err := jwt.HMACCheckHeader(r, secret)

	if err != nil {
		log.Print("Error decoding authentication header ", err)
		return "", err
	}

	if !claims.Valid(time.Now()) {
		log.Print("Authentication header token has expired")
		return "", err
	}

	if claims.Expires.String() == "" {
		log.Print("Authentication header token does not have expiry.")
		return "", nil
	}

	username, ok := claims.String("username")
	if !ok {
		log.Print("username is not present in JWT claims")
		return "", nil
	}

	return username, nil
}
