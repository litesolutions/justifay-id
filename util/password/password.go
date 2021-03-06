package password

import (
	"errors"
	"fmt"
	"os"

	"github.com/trustelem/zxcvbn"
	"golang.org/x/crypto/bcrypt"
)

var (
	MinPasswordLength = 9
	MaxPasswordLength = 72

	// ErrPasswordTooShort ...
	ErrPasswordTooShort = fmt.Errorf(
		"Password must be at least %d characters long",
		MinPasswordLength,
	)

	// ErrPasswordTooLong ...
	ErrPasswordTooLong = fmt.Errorf(
		"Password must be at maximum %d characters long",
		MaxPasswordLength,
	)

	// ErrPasswordTooWeak ...
	ErrPasswordTooWeak = errors.New("Password is too weak")
)

// VerifyPassword compares password and the hashed password
// Fallback to phpass if bcrypt fails
func VerifyPassword(passwordHash, password string) error {
	if bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)) != nil {
		fmt.Fprintln(os.Stderr, "No password match")
		return bcrypt.ErrMismatchedHashAndPassword
	}
	fmt.Fprintln(os.Stderr, "Bcrypt Password matched")
	return nil
}

// HashPassword creates a bcrypt password hash
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 3)
}

// ValidatePassword
func ValidatePassword(password string) error {
	if len(password) < MinPasswordLength {
		return ErrPasswordTooShort
	}

	if len(password) > MaxPasswordLength {
		return ErrPasswordTooLong
	}

	// enforce strong enough passwords
	passwordStrength := zxcvbn.PasswordStrength(password, nil)

	if passwordStrength.Score < 3 {
		return ErrPasswordTooWeak
	}

	return nil
}
