package jwt

import (
	_ "context"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strconv"
)

// VerifyJWT verify JWT token and returns user object
func VerifyJWT(authToken string) (int, string, error) {
	//get jwt secret from environment
	secret := os.Getenv("JWT_SECRET")
	//parse jwt token
	token, err := jwt.ParseWithClaims(authToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return 0, "", err
	}
	uid, err := strconv.Atoi(token.Claims.(*jwt.RegisteredClaims).ID)
	audience := token.Claims.(*jwt.RegisteredClaims).Audience[0]
	return uid, audience, err
}
