package repository

import "github.com/impal-lms/lms-backend/domain"

func (g *GORM) GetAllAlumni() ([]domain.Alumni, error) {
	var Alumni []domain.Alumni
	if err := g.DB.Model(&Alumni).Find(&Alumni).Error; err != nil {
		return Alumni, err
	}

	return Alumni, nil
}

func (g *GORM) CreateAlumni(Alumni domain.Alumni) (domain.Alumni, error) {
	if err := g.DB.Model(&Alumni).Create(&Alumni).Error; err != nil {
		return Alumni, err
	}

	return Alumni, nil
}

func (g *GORM) GetAlumniByID(id int64) (domain.Alumni, error) {
	var Alumni domain.Alumni
	if err := g.DB.Model(&Alumni).Where("id = ?", id).Find(&Alumni).Error; err != nil {
		return Alumni, err
	}

	return Alumni, nil
}

func (g *GORM) UpdateAlumni(Alumni domain.Alumni) error {
	if err := g.DB.Model(&Alumni).Where("id = ?", Alumni.ID).Update(&Alumni).Error; err != nil {
		return err
	}

	return nil
}

func (g *GORM) DeleteAlumniByID(id int64) error {
	var Alumni domain.Alumni
	if err := g.DB.Model(&Alumni).Where("id = ?", id).Delete(&Alumni).Error; err != nil {
		return err
	}

	return nil
}
