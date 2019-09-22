package repository

import "github.com/impal-lms/lms-backend/domain"

type UserRepository interface {
	GetUserByID(id int64) (*domain.User, error)
	Authenticate(email string, password string) (*domain.User, error)
	Save(user domain.User) error
}
