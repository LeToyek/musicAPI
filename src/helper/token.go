package helper

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JWTKey = []byte("secret_key")

type CustomClaims struct {
	ID string `json:"ID"`
	jwt.StandardClaims
}

func GenerateToken(userID string) (string, string, time.Time) {
	expTime := time.Now().Add(time.Minute * 20)
	refExpTime := time.Now().Add(time.Hour * 24)
	claims := &CustomClaims{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	refClaims := &CustomClaims{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refExpTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refClaims)

	tokenStr, err := token.SignedString(JWTKey)
	refTokenStr, err := refToken.SignedString(JWTKey)

	if err != nil {
		panic(err)
	}

	return tokenStr, refTokenStr, expTime
}

func ValidateToken(signedToken string) *CustomClaims {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&CustomClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(JWTKey), nil
		})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		panic(err)
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		panic(err)
	}
	return claims
}
