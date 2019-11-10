package services

import (
	"github.com/impal-lms/lms-backend/authentication"
	"github.com/impal-lms/lms-backend/repository"
)

type LMS struct {
	Repository     repository.Repository
	Authentication authentication.Authentication
}

func NewServices(repo repository.Repository, auth authentication.Authentication) *LMS {
	return &LMS{
		Repository:     repo,
		Authentication: auth,
	}
}
