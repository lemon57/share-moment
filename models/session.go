package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"

	"github.com/lemon57/share-moment/rand"
)

type Session struct {
	ID     int
	UserID int
	// Token is only set when creatting a new session. When lokking up a session
	// this will be left empty, as we only store the hash of a session token
	// in our database and we cannot reverse it into a raw token.
	Token     string
	TokenHash string
}

const (
	MinBytesPerToken = 32
)

type SessionService struct {
	DB *sql.DB
	// BytesPerToken is used to determine how many bytes to use when generating
	// each session token. If this value is not set or is less than the
	// MinBytesPerToken const it will be ignored and MinBytesPerToken will be
	// used.
	BytesPerToken int
}

// Create will create a new session for the user provided. The session token
// will be returned as Token field on the Session type, but only hashed
// session token is stored in the database.
func (ss *SessionService) Create(userID int) (*Session, error) {
	// TODO: Create the session token
	// TODO: Implement SessionService.Create
	bytesPerToken := ss.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}
	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	// TODO: Hash the session token
	session := Session{
		UserID:    userID,
		Token:     token,
		TokenHash: ss.hash(token),
	}
	// TODO: Store the session in our DB
	row := ss.DB.QueryRow(`
		INSERT INTO sessions (user_id, token_hash)
		VALUES ($1, $2) ON CONFLICT (user_id) DO
		UPDATE
		SET token_hash = $2
		RETURNING id;`, session.UserID, session.TokenHash)
	err = row.Scan(&session.ID)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	// TODO: Implement SessionService.User
	// 1. Hash the session token
	tokenHash := ss.hash(token)
	// 2. Query for the session with that hash
	var user User
	row := ss.DB.QueryRow(`
		SELECT users.id,
			users.email,
			users.password_hash
		FROM sessions
			JOIN users ON users.id = sessions.user_id
		WHERE sessions.token_hash = $1`, tokenHash)
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("user: %w", err)
	}
	// 3. return the user
	return &user, nil
}

func (ss *SessionService) Delete(token string) error {
	tokenHash := ss.hash(token)
	_, err := ss.DB.Exec(`
		DELETE FROM sessions
		WHERE token_hash = $1;`, tokenHash)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	return nil
}

func (ss *SessionService) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	// base64 encode the data into a string
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}
