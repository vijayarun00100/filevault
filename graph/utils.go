package graph

import (
	"context"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ctxKey string

const userCtxKey ctxKey = "user_id"

var jwtKey = []byte(processEnv("JWT_CODE"))

func processEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {

		return ""
	}
	return value
}

func generateJWT(userID string) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
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
