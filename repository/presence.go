package repository

import "github.com/impal-lms/lms-backend/domain"

func (g *GORM) GetAllPresence(studentID int64) ([]domain.Presence, error) {
	var Presence []domain.Presence

	if studentID != 0 {
		if err := g.DB.Model(&Presence).Where("student_id = ?", studentID).Find(&Presence).Error; err != nil {
			return Presence, err
		}
	}

	if err := g.DB.Model(&Presence).Find(&Presence).Error; err != nil {
		return Presence, err
	}

	return Presence, nil
}

func (g *GORM) CreatePresence(Presence domain.Presence) (domain.Presence, error) {
	if err := g.DB.Model(&Presence).Create(&Presence).Error; err != nil {
		return Presence, err
	}

	return Presence, nil
}

func (g *GORM) GetPresenceByID(id int64) (domain.Presence, error) {
	var Presence domain.Presence
	if err := g.DB.Model(&Presence).Where("id = ?", id).Find(&Presence).Error; err != nil {
		return Presence, err
	}

	return Presence, nil
}

func (g *GORM) UpdatePresence(Presence domain.Presence) error {
	if err := g.DB.Model(&Presence).Where("id = ?", Presence.ID).Update(&Presence).Error; err != nil {
		return err
	}

	return nil
}

func (g *GORM) DeletePresenceByID(id int64) error {
	var Presence domain.Presence
	if err := g.DB.Model(&Presence).Where("id = ?", id).Delete(&Presence).Error; err != nil {
		return err
	}

	return nil
}
