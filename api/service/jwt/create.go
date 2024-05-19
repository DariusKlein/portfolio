package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strconv"
	"time"
)

func CreateUserJWT(name string, uid int, role string) string {

	//create claims for jwt
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    "portfolio.dariusklein.nl",
		Subject:   name,
		ID:        strconv.Itoa(uid),
		Audience:  []string{role},
	}
	return SignJWT(claims)
}

func SignJWT(claims jwt.Claims) string {
	//Build jwt with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	//get jwt secret from environment
	secret := os.Getenv("JWT_SECRET")
	//sign jwt token with secret
	token, _ := t.SignedString([]byte(secret))
	return token
}
