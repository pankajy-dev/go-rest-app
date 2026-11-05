package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pankajy-dev/go-rest-app/internal/config"
)

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"user":  userId,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(config.JWTSecret()))
}

func VerifyToken(token string) (int64, error) {
	// Using inner anonymous function
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		//Verify the signing method matches with the one we create the token
		// Verifying the type
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.JWTSecret()), nil
	})
	if err != nil {
		return 0, errors.New("could not parse token")
	}

	if !parsedToken.Valid {
		return 0, errors.New("invalid token")
	}

	//Below is how we access claims map

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claim")
	}

	//email := claims["email"].(string)
	userId := int64(claims["user"].(float64))
	return userId, nil
}
