package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)


type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
  db, error := sql.Open("postgres", "postgresql://root:password@localhost:5433/go-chat-db?sslmode=disable")
 if error != nil {
	return nil, error
 }
   
  return &Database{db}, error
}

func (db *Database) Close() error {
	return db.db.Close()
}

func (db *Database) GetDB() *sql.DB {
	return db.db;
}
