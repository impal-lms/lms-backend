package repository

import "github.com/impal-lms/lms-backend/domain"

func (g *GORM) GetAllNotification(userID int64) ([]domain.Notification, error) {
	var Notification []domain.Notification
	if err := g.DB.Model(&Notification).Order("send_at_index DESC").Where("student_id = ? OR teacher_id = ?", userID, userID).Find(&Notification).Error; err != nil {
		return Notification, err
	}

	return Notification, nil
}

func (g *GORM) CreateNotification(Notification domain.Notification) (domain.Notification, error) {
	if err := g.DB.Model(&Notification).Create(&Notification).Error; err != nil {
		return Notification, err
	}

	return Notification, nil
}

func (g *GORM) GetNotificationByID(id int64) (domain.Notification, error) {
	var Notification domain.Notification
	if err := g.DB.Model(&Notification).Where("id = ?", id).Find(&Notification).Error; err != nil {
		return Notification, err
	}

	return Notification, nil
}

func (g *GORM) UpdateNotification(Notification domain.Notification) error {
	if err := g.DB.Model(&Notification).Where("id = ?", Notification.ID).Update(&Notification).Error; err != nil {
		return err
	}

	return nil
}

func (g *GORM) DeleteNotificationByID(id int64) error {
	var Notification domain.Notification
	if err := g.DB.Model(&Notification).Where("id = ?", id).Delete(&Notification).Error; err != nil {
		return err
	}

	return nil
}
