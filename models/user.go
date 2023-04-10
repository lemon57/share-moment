package models

import "database/sql"

type User struct {
	ID           int
	Email        string
	PasswordHash string
}

type UserService struct {
	db *sql.DB
}
