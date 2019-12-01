package repository

import "github.com/impal-lms/lms-backend/domain"

func (g *GORM) GetAllChatRoom(userID int64) ([]domain.ChatRoom, error) {
	var ChatRoom []domain.ChatRoom

	if userID != 0 {
		if err := g.DB.Model(&ChatRoom).Order("last_updated_index DESC").Where("teacher_id = ? OR student_id = ?", userID).Find(&ChatRoom).Error; err != nil {
			return ChatRoom, err
		}
	}

	if err := g.DB.Model(&ChatRoom).Find(&ChatRoom).Error; err != nil {
		return ChatRoom, err
	}

	return ChatRoom, nil
}

func (g *GORM) CreateChatRoom(ChatRoom domain.ChatRoom) (domain.ChatRoom, error) {
	if err := g.DB.Model(&ChatRoom).Create(&ChatRoom).Error; err != nil {
		return ChatRoom, err
	}

	return ChatRoom, nil
}

func (g *GORM) GetChatRoomByID(id int64) (domain.ChatRoom, error) {
	var ChatRoom domain.ChatRoom
	if err := g.DB.Model(&ChatRoom).Where("id = ?", id).Find(&ChatRoom).Error; err != nil {
		return ChatRoom, err
	}

	return ChatRoom, nil
}

func (g *GORM) UpdateChatRoom(ChatRoom domain.ChatRoom) error {
	if err := g.DB.Model(&ChatRoom).Where("id = ?", ChatRoom.ID).Update(&ChatRoom).Error; err != nil {
		return err
	}

	return nil
}

func (g *GORM) DeleteChatRoomByID(id int64) error {
	var ChatRoom domain.ChatRoom
	if err := g.DB.Model(&ChatRoom).Where("id = ?", id).Delete(&ChatRoom).Error; err != nil {
		return err
	}

	return nil
}

func (g *GORM) GetChatRoomBetweenUsers(studentID, teacherID int64) (domain.ChatRoom, error) {
	var ChatRoom domain.ChatRoom
	if err := g.DB.Model(&ChatRoom).Order("last_updated_index DESC").Where("student_id = ? AND teacher_id = ?", studentID, teacherID).Find(&ChatRoom).Error; err != nil {
		return ChatRoom, err
	}

	return ChatRoom, nil
}

func (g *GORM) GetAllMessage(ChatRoomID int64) ([]domain.Message, error) {
	var Message []domain.Message

	if ChatRoomID != 0 {
		if err := g.DB.Model(&Message).Order("send_at_index DESC").Where("chat_room_id = ?", ChatRoomID).Find(&Message).Error; err != nil {
			return Message, err
		}
	}

	if err := g.DB.Model(&Message).Find(&Message).Error; err != nil {
		return Message, err
	}

	return Message, nil
}

func (g *GORM) CreateMessage(Message domain.Message) (domain.Message, error) {
	if err := g.DB.Model(&Message).Create(&Message).Error; err != nil {
		return Message, err
	}

	return Message, nil
}

func (g *GORM) GetMessageByID(id int64) (domain.Message, error) {
	var Message domain.Message
	if err := g.DB.Model(&Message).Where("id = ?", id).Find(&Message).Error; err != nil {
		return Message, err
	}

	return Message, nil
}

func (g *GORM) UpdateMessage(Message domain.Message) error {
	if err := g.DB.Model(&Message).Where("id = ?", Message.ID).Update(&Message).Error; err != nil {
		return err
	}

	return nil
}

func (g *GORM) DeleteMessageByID(id int64) error {
	var Message domain.Message
	if err := g.DB.Model(&Message).Where("id = ?", id).Delete(&Message).Error; err != nil {
		return err
	}

	return nil
}
