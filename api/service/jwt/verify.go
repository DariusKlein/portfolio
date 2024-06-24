package jwt

import (
	_ "context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strconv"
)

func VerifyUser(r *http.Request) (int, string, error) {
	bearerToken := r.Header.Get("Authorization")
	jwtToken := ""
	if len(bearerToken) > 7 {
		jwtToken = bearerToken[len("Bearer "):]
	} else {
		return 0, "", fmt.Errorf("token empty")
	}
	return VerifyJWT(jwtToken)
}

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
