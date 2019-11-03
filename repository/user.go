package repository

import (
	"github.com/impal-lms/lms-backend/domain"
)

func (g *GORM) GetAllUser() ([]domain.User, error) {
	var user []domain.User

	if err := g.DB.Model(&user).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (g *GORM) CreateUser(user domain.User) (domain.User, error) {
	if err := g.DB.Model(&user).Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (g *GORM) GetUserByID(id int64) (domain.User, error) {
	var user domain.User

	if err := g.DB.Model(&user).Where("id = ?", id).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (g *GORM) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User

	if err := g.DB.Model(&user).Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (g *GORM) UpdateUser(user domain.User) error {
	if err := g.DB.Model(&user).Where("id = ?", user.ID).Update(&user).Error; err != nil {
		return err
	}

	return nil
}

func (g *GORM) DeleteUserById(id int64) error {
	var user domain.User
	if err := g.DB.Model(&user).Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
