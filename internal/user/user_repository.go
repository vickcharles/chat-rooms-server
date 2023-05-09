package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastInsertID int64
	_, err := r.db.QueryContext(ctx, "INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, user.Password)

	if err != nil {
		return nil, err
	}

	  user.ID = int64(lastInsertID)
	  return user, nil
	
}

func (repository *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
     u := User{}

	 query := "SELECT * FROM users WHERE email = $1"

	 err := repository.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Username, &u.Email, &u.Password)

	 if err != nil {
		 return &User{}, err
	 }

	 return &u, nil
}
