package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/quocquann/book-api/app/queries"
)

type Queries struct {
	*queries.BookQueries
}

func OpenConnection() (*Queries, error) {
	cfg := mysql.Config{
		Net:    "tcp",
		User:   os.Getenv("DB_USERNAME"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Addr:   os.Getenv("DB_ADDRESS"),
		DBName: os.Getenv("DB_NAME"),
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connect to DB successfully")
	return &Queries{
		BookQueries: &queries.BookQueries{DB: db},
	}, nil
}
