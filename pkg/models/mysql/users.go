package mysql

import (
	"database/sql"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/go-sql-driver/mysql"
	"github.com/sk000f/snippetbox/pkg/models"
)

// UserModel represents an authenticated user
type UserModel struct {
	DB *sql.DB
}

// Insert adds a new user to the database
func (m *UserModel) Insert(name, email, password string) error {

	// Create bcrypt hash of password from signup form
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, hashed_password, created) VALUES(?, ?, ?, UTC_TIMESTAMP())`

	_, err = m.DB.Exec(stmt, name, email, string(hashedPassword))
	if err != nil {
		var mySQLError *mysql.MySQLError
		if errors.As(err, &mySQLError) {
			if mySQLError.Number == 1062 && strings.Contains(mySQLError.Message, "users_uc_email") {
				return models.ErrDuplicateEmail
			}
		}
		return err
	}

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
