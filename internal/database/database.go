package database

import (
	"database/sql"
	"vpn-handler/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	DB *sql.DB
}

func NewDB() (*DB, error) {
	var err error
	db := &DB{}

	db.DB, err = sql.Open("sqlite3", "connection")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *DB) Close() {
	d.DB.Close()
}

func (d *DB) CreateUser(name, email string) error {
	return nil
}

func (d *DB) ReadUser(id int) (*models.User, error) {
	return nil, nil
}

func (d *DB) UpdateUser(id int, name, email string) error {
	return nil
}

func (d *DB) DeleteUser(id int) error {
	return nil
}
