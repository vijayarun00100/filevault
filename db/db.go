package db

import (
    "context"
    "github.com/jackc/pgx/v5/pgxpool"
)

// DB wraps the connection pool
type DB struct {
    Conn *pgxpool.Pool
}

// NewDB connects to PostgreSQL and returns a DB instance
func NewDB(connString string) (*DB, error) {
    pool, err := pgxpool.New(context.Background(), connString)
    if err != nil {
        return nil, err
    }
    return &DB{Conn: pool}, nil
}
