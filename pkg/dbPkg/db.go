package dbPkg

import (
	"Chat-Websocket/config"
	"Chat-Websocket/internal/db"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Database struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewDatabase(cfg *config.Config) (*Database, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Postgres.PostgresqlUser,
		cfg.Postgres.PostgresqlPassword,
		cfg.Postgres.PostgresqlHost,
		cfg.Postgres.PostgresqlPort,
		cfg.Postgres.PostgresqlDbname,
		cfg.Postgres.PostgresqlSSLMode,
	)

	dbConn, err := sql.Open(cfg.Postgres.PgDriver, dsn)
	if err != nil {
		return nil, err
	}

	queries := db.New(dbConn)
	return &Database{DB: dbConn, Queries: queries}, nil
}

func (d *Database) Close() {
	d.DB.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.DB
}
