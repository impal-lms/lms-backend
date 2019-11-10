package repository

import "github.com/impal-lms/lms-backend/domain"

func (g *GORM) GetAllClassroom() ([]domain.Classroom, error) {
	var classroom []domain.Classroom
	if err := g.DB.Model(&classroom).Find(&classroom).Error; err != nil {
		return classroom, err
	}

	return classroom, nil
}

func (g *GORM) CreateClassroom(classroom domain.Classroom) (domain.Classroom, error) {
	if err := g.DB.Model(&classroom).Create(&classroom).Error; err != nil {
		return classroom, err
	}

	return classroom, nil
}

func (g *GORM) GetClassroomByID(id int64) (domain.Classroom, error) {
	var classroom domain.Classroom
	if err := g.DB.Model(&classroom).Where("id = ?", id).Find(&classroom).Error; err != nil {
		return classroom, err
	}

	return classroom, nil
}

func (g *GORM) UpdateClassroom(classroom domain.Classroom) error {
	if err := g.DB.Model(&classroom).Where("id = ?", classroom.ID).Update(&classroom).Error; err != nil {
		return err
	}

	return nil
}

func (g *GORM) DeleteClassroomById(id int64) error {
	var classroom domain.Classroom
	if err := g.DB.Model(&classroom).Where("id = ?", id).Delete(&classroom).Error; err != nil {
		return err
	}

	return nil
}
