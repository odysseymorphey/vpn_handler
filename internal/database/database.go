package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	DB *sql.DB
}

func NewDB() (*DB, error) {
	var err error
	db := &DB{}

	db.DB, err = sql.Open("sqlite", "connection")
	if err != nil {
		return nil, err
	}
	
	return db, nil
}

func (d *DB) Close() {
    d.DB.Close()
}
