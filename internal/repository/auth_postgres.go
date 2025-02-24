package repository

import (
	fin "FinTransaction"
	"database/sql"
	"fmt"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthDB(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user fin.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password) VALUES ($1, $2, $3) RETURNING id", usersTable)
	err := r.db.QueryRow(query, user.Name, user.Username, user.Password).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (fin.User, error) {
	var user fin.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE username = $1 AND password = $2", usersTable)
	err := r.db.QueryRow(query, username, password).Scan(&user.ID, &user.Name, &user.Username, &user.Password)

	return user, err
}
