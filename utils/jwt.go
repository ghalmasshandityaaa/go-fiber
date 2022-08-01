package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
)

var JWT_SECRET_KEY string = "apaajayangpentingaman"

func GenerateAccessToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	webToken, err := token.SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}

	return webToken, nil
}

func VerifyAccessToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(JWT_SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func DecodeToken(jwtToken string) (jwt.MapClaims, error) {
	token, err := VerifyAccessToken(jwtToken)
	if err != nil {
		return nil, err
	}

	claims, isValid := token.Claims.(jwt.MapClaims)
	if isValid && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
