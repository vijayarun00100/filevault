package graph

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ctxKey string

const userCtxKey ctxKey = "user_id"


func getJWTSecret() string {
	s := "dev-secret"
	if env := os.Getenv("JWT_CODE"); env != "" {
		s = env
	}
	return s
}

func generateJWT(userID string) (string, error) {
	jwtKey := []byte(getJWTSecret())
	log.Printf("Generating JWT with secret: %s", string(jwtKey))

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetUserCtxKey() ctxKey {
	return userCtxKey
}

func GetUserIDFromCtx(ctx context.Context) (string, bool) {
	uid, ok := ctx.Value(userCtxKey).(string)
	return uid, ok
}
