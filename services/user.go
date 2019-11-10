package services

import (
	"time"

	"github.com/impal-lms/lms-backend/domain"
	"github.com/pkg/errors"
)

func (lms *LMS) GetAllUser() ([]domain.User, error) {
	return lms.Repository.GetAllUser()
}

func (lms *LMS) CreateUser(user domain.User) (domain.User, int, error) {
	hash, err := lms.Authentication.HashPassword(user.Password)
	if err != nil {
		return domain.User{}, 400, err
	}

	user.ID = time.Now().UnixNano()
	user.Role = -1
	user.Password = hash

	user, err = lms.Repository.CreateUser(user)
	if err != nil {
		return domain.User{}, 400, err
	}

	return user, 201, nil
}

func (lms *LMS) GetUserByID(id int64) (domain.User, error) {
	return lms.Repository.GetUserByID(id)
}

func (lms *LMS) GetUserByEmail(email string) (domain.User, error) {
	return lms.Repository.GetUserByEmail(email)
}

func (lms *LMS) UpdateUser(user domain.User) (domain.User, error) {
	user, err := lms.Repository.GetUserByID(user.ID)
	if err != nil {
		return domain.User{}, errors.Wrap(err, "user does not exist")
	}

	return user, lms.Repository.UpdateUser(user)
}

func (lms *LMS) DeleteUserById(id int64) (domain.User, error) {
	user, err := lms.Repository.GetUserByID(id)
	if err != nil {
		return domain.User{}, errors.Wrap(err, "user does not exist")
	}

	return domain.User{}, lms.Repository.DeleteUserById(user.ID)
}
