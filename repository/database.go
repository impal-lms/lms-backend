package repository

import (
	"github.com/impal-lms/lms-backend/domain"
	"github.com/jinzhu/gorm"
)

type GORM struct {
	DB *gorm.DB
}

func NewRepository(DB *gorm.DB) *GORM {
	DB.AutoMigrate(
		&domain.User{},
		&domain.Material{},
		&domain.Task{},
		&domain.Submission{},
	)

	return &GORM{
		DB: DB,
	}
}
