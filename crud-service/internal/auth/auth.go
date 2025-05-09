package auth

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// ParseJWT verifies the token and returns the claims.
func ParseJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err // error parsing (invalid token format, bad signature, etc.)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
