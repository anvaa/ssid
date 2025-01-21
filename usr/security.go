package users

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
)

var (
	emailRegex    = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	ErrInvalidEmail = errors.New("not a valid email")
	ErrEmptyPassword = errors.New("password can't be empty")
	ErrShortPassword = errors.New("password must be at least 8 characters")
	ErrLongPassword  = errors.New("password must be less than 255 characters")
)

func IsValidEmail(email string) error {
	if !emailRegex.MatchString(email) {
		return ErrInvalidEmail
	}
	return nil
}

func IsValidPassword(password string) error {
	switch {
	case password == "":
		return ErrEmptyPassword
	case len(password) < 8:
		return ErrShortPassword
	case len(password) > 255:
		return ErrLongPassword
	default:
		return nil
	}
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewUuid() int {
	return int(uuid.New().ID())
}
