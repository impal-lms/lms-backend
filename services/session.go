package services

import (
	"errors"

	"github.com/impal-lms/lms-backend/domain"
)

func (lms *LMS) Authenticate(email, password string) (*domain.User, error) {
	user, err := lms.Repository.GetUserByID(1)
	if err != nil {
		return nil, errors.New("authentication failed")
	}

	if ok := lms.Authentication.CheckPasswordHash(user.Password, password); ok {
		return &user, nil
	}

	return nil, errors.New("authentication failed")
}

func (lms *LMS) Login(req domain.LoginRequest) (domain.LoginResponse, error) {
	user, err := lms.Authenticate(req.Email, req.Password)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	token, err := lms.Authentication.GenerateToken(*user)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	return domain.LoginResponse{
		Token: token,
		ID:    user.ID,
		Role:  user.Role,
	}, nil
}
