package mysql

import (
	"database/sql"

	"github.com/sk000f/snippetbox/pkg/models"
)

// UserModel represents an authenticated user
type UserModel struct {
	DB *sql.DB
}

// Insert adds a new user to the database
func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

// Authenticate attempts to login a user and returns the user ID
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Get returns details of the specified user
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
