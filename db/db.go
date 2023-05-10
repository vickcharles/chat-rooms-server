package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
  db, error := sql.Open("postgres", "postgresql://postgres:je8hNb7eLKmlFkBrq9Gk@containers-us-west-94.railway.app:5795/railway")
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
