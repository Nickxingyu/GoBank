package jwt

import (
	"fmt"

	"github.com/Nickxingyu/GoBank/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type AuthClaim struct {
	jwt.RegisteredClaims
	UserID uint `json:"user_id"`
}

var jwtKey = config.Config.JWT.Key

func GenerateTokenString(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*AuthClaim, error) {
	authClaim := AuthClaim{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		&authClaim,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("jwt invalid")
	}

	if _, ok := token.Claims.(AuthClaim); !ok {
		return nil, fmt.Errorf("unknown claim type")
	}

	return &authClaim, nil
}
