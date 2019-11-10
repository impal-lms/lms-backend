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
	_, err := lms.Repository.GetUserByID(user.ID)
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

func (lms *LMS) ChangePassword(user domain.User, req domain.ChangePasswordRequest) (domain.User, error) {
	userVerify, err := lms.Repository.GetUserByID(user.ID)
	if err != nil {
		return domain.User{}, errors.Wrap(err, "user does not exist")
	}

	if ok := lms.Authentication.CheckPasswordHash(userVerify.Password, req.OldPassword); !ok {
		return domain.User{}, errors.New("old password does not match")
	}

	if req.OldPassword == req.NewPassword {
		return domain.User{}, errors.New("new password can not be same as old password")
	}

	hash, err := lms.Authentication.HashPassword(req.NewPassword)
	if err != nil {
		return domain.User{}, errors.New("change password error")
	}

	user.Password = hash

	return user, lms.Repository.UpdateUser(user)
}

func (lms *LMS) ChangeRole(user domain.User, req domain.ChangeRoleRequest) (domain.User, error) {
	_, err := lms.Repository.GetUserByID(user.ID)
	if err != nil {
		return domain.User{}, errors.Wrap(err, "user does not exist")
	}

	user.Role = req.Role

	return user, lms.Repository.UpdateUser(user)
}
