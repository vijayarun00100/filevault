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
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	connString := "postgres://postgres:root@localhost:5432/filevault?sslmode=disable"
	database, err := db.NewDB(connString)
	if err != nil {
		log.Fatal(err)
	}

	resolver := &graph.Resolver{DB: database}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
	srv.AddTransport(transport.MultipartForm{})
	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", AuthMiddleware(srv))

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		auth := r.Header.Get("Authorization")
		if auth == "" {
			log.Println("No Authorization header found")
			next.ServeHTTP(w, r)
			return
		}

		// Expect "Bearer <token>"
		tokenStr := strings.TrimPrefix(auth, "Bearer ")
		tokenStr = strings.TrimSpace(tokenStr)
		if tokenStr == "" {
			log.Println("Empty token after Bearer prefix")
			next.ServeHTTP(w, r)
			return
		}

		log.Printf("Processing token: %s...", tokenStr[:20])

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(getJWTSecret()), nil
		})
		if err != nil {
			log.Printf("JWT parsing error: %v", err)
		} else if !token.Valid {
			log.Println("JWT token is not valid")
		} else {
			if uid, ok := claims["user_id"].(string); ok {
				log.Printf("Successfully authenticated user: %s", uid)
				ctx := context.WithValue(r.Context(), graph.GetUserCtxKey(), uid)
				r = r.WithContext(ctx)
			} else {
				log.Println("No user_id found in claims")
			}
		}
		next.ServeHTTP(w, r)
	})
}

func getJWTSecret() string {
	return os.Getenv("JWT_CODE")
}
