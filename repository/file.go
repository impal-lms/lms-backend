package repository

import (
	"github.com/impal-lms/lms-backend/domain"
)

func (g *GORM) GetAllMaterial() ([]domain.Material, error) {
	var Material []domain.Material

	if err := g.DB.Model(&Material).Find(&Material).Error; err != nil {
		return Material, err
	}

	return Material, nil
}

func (g *GORM) CreateMaterial(Material domain.Material) (domain.Material, error) {
	if err := g.DB.Model(&Material).Create(&Material).Error; err != nil {
		return Material, err
	}

	return Material, nil
}

func (g *GORM) GetMaterialByID(id int64) (domain.Material, error) {
	var Material domain.Material

	if err := g.DB.Model(&Material).Where("id = ?", id).Find(&Material).Error; err != nil {
		return Material, err
	}

	return Material, nil
}

func (g *GORM) UpdateMaterial(Material domain.Material) error {
	if err := g.DB.Model(&Material).Where("id = ?", Material.ID).Update(&Material).Error; err != nil {
		return err
	}

	return nil
}

func (g *GORM) DeleteMaterialById(id int64) error {
	var Material domain.Material
	if err := g.DB.Model(&Material).Where("id = ?", id).Delete(&Material).Error; err != nil {
		return err
	}

	return nil
}

func (g *GORM) GetAllTask() ([]domain.Task, error) {
	var Task []domain.Task

	if err := g.DB.Model(&Task).Find(&Task).Error; err != nil {
		return Task, err
	}

	return Task, nil
}

func (g *GORM) CreateTask(Task domain.Task) (domain.Task, error) {
	if err := g.DB.Model(&Task).Create(&Task).Error; err != nil {
		return Task, err
	}

	return Task, nil
}

func (g *GORM) GetTaskByID(id int64) (domain.Task, error) {
	var Task domain.Task

	if err := g.DB.Model(&Task).Where("id = ?", id).Find(&Task).Error; err != nil {
		return Task, err
	}

	return Task, nil
}

func (g *GORM) UpdateTask(Task domain.Task) error {
	if err := g.DB.Model(&Task).Where("id = ?", Task.ID).Update(&Task).Error; err != nil {
		return err
	}

	return nil
}

func (g *GORM) DeleteTaskById(id int64) error {
	var Task domain.Task
	if err := g.DB.Model(&Task).Where("id = ?", id).Delete(&Task).Error; err != nil {
		return err
	}

	return nil
}

func (g *GORM) GetAllSubmission() ([]domain.Submission, error) {
	var Submission []domain.Submission

	if err := g.DB.Model(&Submission).Find(&Submission).Error; err != nil {
		return Submission, err
	}

	return Submission, nil
}

func (g *GORM) CreateSubmission(Submission domain.Submission) (domain.Submission, error) {
	if err := g.DB.Model(&Submission).Create(&Submission).Error; err != nil {
		return Submission, err
	}

	return Submission, nil
}

func (g *GORM) GetSubmissionByID(id int64) (domain.Submission, error) {
	var Submission domain.Submission

	if err := g.DB.Model(&Submission).Where("id = ?", id).Find(&Submission).Error; err != nil {
		return Submission, err
	}

	return Submission, nil
}

func (g *GORM) UpdateSubmission(Submission domain.Submission) error {
	if err := g.DB.Model(&Submission).Where("id = ?", Submission.ID).Update(&Submission).Error; err != nil {
		return err
	}

	return nil
}

func (g *GORM) DeleteSubmissionById(id int64) error {
	var Submission domain.Submission
	if err := g.DB.Model(&Submission).Where("id = ?", id).Delete(&Submission).Error; err != nil {
		return err
	}

	return nil
}
