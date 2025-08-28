package bootstrap

import (
	"database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

func NewPostgresDatabase(env *Env) *sql.DB {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", env.DBUser, env.DBPass, env.DBHost, env.DBPort, env.DBName)
	db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }

	return db
}