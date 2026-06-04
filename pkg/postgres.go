package pkg

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sql.DB
}

func InitPostgres() (*Postgres, error) {
	connSrt := "user=postgres dbname=postgres password=pguhbmi@1 host=localhost port=5433 sslmode=disable"

	db, err := sql.Open("postgres", connSrt)
	if err != nil {
		return nil, err
	}
	return &Postgres{DB: db}, nil
}
