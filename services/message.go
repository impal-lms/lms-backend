package services

import (
	"time"

	"github.com/impal-lms/lms-backend/domain"
	"github.com/pkg/errors"
)

func (lms *LMS) GetAllChatRoom(userID int64) ([]domain.ChatRoom, error) {
	return lms.Repository.GetAllChatRoom(userID)
}

func (lms *LMS) GetChatRoomByID(ChatRoomID int64) (domain.ChatRoom, error) {
	return lms.Repository.GetChatRoomByID(ChatRoomID)
}

func (lms *LMS) CreateChatRoom(ChatRoom domain.ChatRoom) (domain.ChatRoom, int, error) {
	ChatRoomCheck, _ := lms.Repository.GetChatRoomBetweenUsers(ChatRoom.StudentID, ChatRoom.TeacherID)
	if ChatRoom.StudentID == ChatRoomCheck.StudentID && ChatRoom.TeacherID == ChatRoomCheck.TeacherID {
		return domain.ChatRoom{}, 400, errors.New("ChatRoom is already existed")
	}

	ChatRoom.ID = time.Now().UnixNano()
	ChatRoom.LastUpdated = time.Now()
	ChatRoom.LastUpdatedIndex = ChatRoom.LastUpdated.UnixNano()

	ChatRoom, err := lms.Repository.CreateChatRoom(ChatRoom)
	if err != nil {
		return domain.ChatRoom{}, 400, err
	}

	return ChatRoom, 201, nil

}

func (lms *LMS) UpdateChatRoom(ChatRoom domain.ChatRoom) (domain.ChatRoom, error) {
	_, err := lms.Repository.GetChatRoomByID(ChatRoom.ID)
	if err != nil {
		return domain.ChatRoom{}, errors.Wrap(err, "ChatRoom does not exist")
	}

	return ChatRoom, lms.Repository.UpdateChatRoom(ChatRoom)
}

func (lms *LMS) DeleteChatRoomByID(ChatRoomID int64) (domain.ChatRoom, error) {
	ChatRoom, err := lms.Repository.GetChatRoomByID(ChatRoomID)
	if err != nil {
		return domain.ChatRoom{}, errors.Wrap(err, "ChatRoom does not exist")
	}

	return ChatRoom, lms.Repository.DeleteChatRoomByID(ChatRoom.ID)
}

func (lms *LMS) GetAllMessage(ChatRoomID int64) ([]domain.Message, error) {
	return lms.Repository.GetAllMessage(ChatRoomID)
}

func (lms *LMS) GetMessageByID(messageID int64) (domain.Message, error) {
	return lms.Repository.GetMessageByID(messageID)
}

func (lms *LMS) CreateMessage(message domain.Message) (domain.Message, int, error) {
	message.ID = time.Now().UnixNano()
	message.SendAt = time.Now()
	message.SendAtIndex = message.SendAt.UnixNano()

	message, err := lms.Repository.CreateMessage(message)
	if err != nil {
		return domain.Message{}, 400, err
	}

	ChatRoom, err := lms.GetChatRoomByID(message.ChatRoomID)
	if err != nil {
		return domain.Message{}, 400, err
	}

	ChatRoom.LastUpdated = message.SendAt
	ChatRoom.LastUpdatedIndex = ChatRoom.LastUpdated.UnixNano()

	ChatRoom, err = lms.UpdateChatRoom(ChatRoom)
	if err != nil {
		return domain.Message{}, 400, err
	}

	var notification domain.Notification
	notification.ReceiverID = message.ReceiverID
	notification.SendAt = message.SendAt
	notification.SendAtIndex = message.SendAtIndex
	notification.Content = message.Content

	notification, _, err = lms.CreateNotification(notification)
	if err != nil {
		return domain.Message{}, 400, err
	}

	return message, 201, nil
}

func (lms *LMS) UpdateMessage(message domain.Message) (domain.Message, error) {
	_, err := lms.Repository.GetMessageByID(message.ID)
	if err != nil {
		return domain.Message{}, errors.Wrap(err, "Message does not exist")
	}

	return message, lms.Repository.UpdateMessage(message)
}

func (lms *LMS) DeleteMessageByID(messageID int64) (domain.Message, error) {
	message, err := lms.Repository.GetMessageByID(messageID)
	if err != nil {
		return domain.Message{}, errors.Wrap(err, "Message does not exist")
	}

	return message, lms.Repository.DeleteMessageByID(messageID)
}
