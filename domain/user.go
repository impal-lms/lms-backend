package domain

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserRole int64

const (
	Admin   UserRole = 0
	Teacher UserRole = 1
	Student UserRole = 2
)

type User struct {
	Id       int64
	Name     string
	Email    string
	Password string
	Role     UserRole
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (u User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func New(name, email, password string, role UserRole) (*User, error) {

	pass, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		Id:       time.Now().UnixNano(),
		Name:     name,
		Email:    email,
		Password: pass,
	}, nil
}
