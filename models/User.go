package models

import (
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"hashed_password"`
}

type UserModel struct {
	DB *sqlx.DB
}

func NewUserModel(db *sqlx.DB) *UserModel {
	return &UserModel{DB: db}
}

func (m *UserModel) Register(username, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (username, email, hashed_password) VALUES (:username, :email, :hashed_password)`
	_, err = m.DB.NamedExec(query, map[string]interface{}{
		"username":        username,
		"email":           email,
		"hashed_password": string(hashedPassword),
	})

	return err
}
