package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"filevault/db"
	"filevault/graph"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		connString = "postgres://postgres:root@localhost:5432/filevault?sslmode=disable"
	}
	database, err := db.NewDB(connString)
	if err != nil {
		log.Fatal(err)
	}

	graph.InitSupabase()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: database}}))
	srv.AddTransport(transport.Websocket{})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	// Add error handling
	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		log.Printf("GraphQL Error: %v", e)
		return graphql.DefaultErrorPresenter(ctx, e)
	})

	// Enable CORS for frontend
	corsHandler := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Max-Age", "86400")
			
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			
			next.ServeHTTP(w, r)
		})
	}

	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", corsHandler(AuthMiddleware(srv)))
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("./uploads"))))
	http.HandleFunc("/download/", PublicDownloadHandler(database))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		log.Printf("Content-Type: %s", r.Header.Get("Content-Type"))

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

		tokenPreview := tokenStr
		if len(tokenStr) > 20 {
			tokenPreview = tokenStr[:20] + "..."
		}
		log.Printf("Processing token: %s", tokenPreview)

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

func PublicDownloadHandler(database *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract file ID from URL path
		fileID := strings.TrimPrefix(r.URL.Path, "/download/")
		if fileID == "" {
			http.Error(w, "File ID required", http.StatusBadRequest)
			return
		}

		log.Printf("Public download request for file ID: %s", fileID)

		// Get file info from database
		var filename, path string
		err := database.Conn.QueryRow(context.Background(),
			"SELECT filename, path FROM files WHERE id=$1", fileID).Scan(&filename, &path)

		if err != nil {
			log.Printf("File not found: %v", err)
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}

		bucket := "filevault"
		supabaseURL := os.Getenv("SUPABASE_URL")
		downloadURL := fmt.Sprintf("%s/storage/v1/object/public/%s/%s", supabaseURL, bucket, path)

		log.Printf("Redirecting to: %s", downloadURL)

		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))

		http.Redirect(w, r, downloadURL, http.StatusFound)
	}
}
