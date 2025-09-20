package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"filevault/db"
	"filevault/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
	connString := "postgres://postgres:root@localhost:5432/filevault?sslmode=disable"
	database, err := db.NewDB(connString)
	if err != nil {
		log.Fatal(err)
	}

	resolver := &graph.Resolver{DB: database}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", AuthMiddleware(srv))

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		auth := r.Header.Get("Authorization")
		if auth == "" {
			next.ServeHTTP(w, r)
			return
		}

		// Expect "Bearer <token>"
		tokenStr := strings.TrimPrefix(auth, "Bearer ")
		tokenStr = strings.TrimSpace(tokenStr)
		if tokenStr == "" {
			next.ServeHTTP(w, r)
			return
		}

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(getJWTSecret()), nil
		})
		if err == nil && token.Valid {
			if uid, ok := claims["user_id"].(string); ok {
				ctx := context.WithValue(r.Context(), graph.GetUserCtxKey(), uid)
				r = r.WithContext(ctx)
			}
		}
		next.ServeHTTP(w, r)
	})
}

func getJWTSecret() string {
	s := "dev-secret"
	if env := os.Getenv("JWT_CODE"); env != "" {
		s = env
	}
	return s
}
