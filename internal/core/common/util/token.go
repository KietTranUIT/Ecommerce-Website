package util

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (bool, error, string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return false, err, ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Trích xuất thông tin từ JWT
		username := claims["username"].(string)
		log.Println("JWT: ", username)
		return true, nil, username
	}

	return false, nil, ""
}
