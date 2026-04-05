package repository

import (
	"database/sql"
	"errors"

	"github.com/r200a/vc-platform/internal/auth/model"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(email, passwordHash, role string) (string, error) {
	var userID string
	err := r.db.QueryRow(
		`INSERT INTO users (email, password_hash, role)
         VALUES ($1, $2, $3)
         RETURNING user_id`,
		email, passwordHash, role,
	).Scan(&userID)
	return userID, err
}

func (r *AuthRepository) GetUserByEmail(email string) (model.User, error) {
	var u model.User
	err := r.db.QueryRow(
		`SELECT user_id, email, password_hash, role, is_active
         FROM users WHERE email = $1`,
		email,
	).Scan(&u.UserID, &u.Email, &u.PasswordHash, &u.Role, &u.IsActive)
	if err == sql.ErrNoRows {
		return u, errors.New("user not found")
	}
	return u, err
}

func (r *AuthRepository) SaveRefreshToken(userID, token string) error {
	_, err := r.db.Exec(
		`UPDATE users SET refresh_token = $1, updated_at = NOW()
         WHERE user_id = $2`,
		token, userID,
	)
	return err
}

func (r *AuthRepository) GetUserByRefreshToken(token string) (model.User, error) {
	var u model.User
	err := r.db.QueryRow(
		`SELECT user_id, email, role FROM users WHERE refresh_token = $1`,
		token,
	).Scan(&u.UserID, &u.Email, &u.Role)
	if err == sql.ErrNoRows {
		return u, errors.New("invalid refresh token")
	}
	return u, err
}
