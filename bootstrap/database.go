package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/stdlib"
)

const (
	driverName = "pgx"
)

func NewPostgresDatabase(env *Env) *sqlx.DB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s", env.DBUser, env.DBName, env.DBPass, env.DBHost)
	db, err := sqlx.ConnectContext(ctx, driverName, connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
