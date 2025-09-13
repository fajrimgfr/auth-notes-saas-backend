package util

import (
	"time"

	"github.com/fajrimgfr/auth-notes-saas-backend/domain"
	"github.com/golang-jwt/jwt/v4"
)

func CreateAccessToken(user *domain.User, expiry int, secret string) (string, error) {
	exp := jwt.NumericDate{Time: time.Now().Add(time.Hour * time.Duration(expiry))}
	claims := &domain.JwtCustomClaim{
		Name: user.Name.String,
		ID:   user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func CreateRefreshToken(user *domain.User, expiry int, secret string) (string, error) {
	exp := jwt.NumericDate{Time: time.Now().Add(time.Hour * time.Duration(expiry))}
	claims := &domain.JwtCustomClaim{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func IsAuthorized() {

}
