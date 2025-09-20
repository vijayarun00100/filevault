package main


import (
    "log"
    "net/http"

    "filevault/db"
    "filevault/graph"
    
    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"
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
    http.Handle("/query", srv)

    log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
