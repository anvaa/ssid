package users

import (
	"errors"
	"regexp"
	
	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
)

func IsValidEmail(email string) error {
	var err error

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		err = errors.New("not a valid email")
		return err
	}

	return nil
}

func IsValidPassword(password string) error {
	var err error

	if password == "" {
		err = errors.New("password can't be empty")
		return err
	}

	if len(password) < 8 {
		err = errors.New("password must be at least 8 characters")
		return err
	}

	if len(password) > 50 {
		err = errors.New("password must be less than 50 characters")
		return err
	}

	return nil
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
	// generste a new uuid
	return int(uuid.New().ID())
}