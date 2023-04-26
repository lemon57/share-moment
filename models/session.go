package models

import "database/sql"

type Session struct {
	ID        int
	UserId    int
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	// TODO: Create the session token
	// TODO: Implement SessionService.Create
	return nil, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	// TODO: Implement SessionService.User
	return nil, nil
}
