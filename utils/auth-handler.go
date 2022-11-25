package utils

import (
	"fmt"
	"time"

	"github.com/shadowshot-x/micro-product-go/authservice/jwt"
)

func GetSignedToken() (string, error) {

	claimsMap := jwt.ClaimsMap{
		Aud: "frontend",
		Iss: "get-working",
		Exp: fmt.Sprint(time.Now().Add(time.Minute * 1).Unix()),
	}

	secret := "Secure_Random_String"
	header := "HS256"
	tokenString, err := jwt.GenerateToken(header, claimsMap, secret)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}
