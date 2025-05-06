package utils

import (
	"portfolio-user-service/repository/auth/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWTWithUserDetails generates a JWT token with user information
func GenerateJWT(user models.User) (string, error) {
	secretKey := LoadEnvVar("JWT_SECRET")
	tokenExpiryTime := time.Now().Add(time.Hour * 24).Unix()
	if LoadEnvVar("SERVER") == "DEVELOPMENT" {
		tokenExpiryTime = time.Now().Add(time.Hour * 1).Unix()
	}

	claims := jwt.MapClaims{
		"id":    user.ID,
		"name":  user.FullName,
		"email": user.Email,
		"exp":   tokenExpiryTime,
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
