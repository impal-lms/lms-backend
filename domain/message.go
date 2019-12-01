package domain

import "time"

type ChatRoom struct {
	ID               int64     `json:"id" gorm:"primary_key"`
	TeacherID        int64     `json:"teacher_id"`
	StudentID        int64     `json:"student_id"`
	LastUpdated      time.Time `json:"last_updated"`
	LastUpdatedIndex int64     `json:"-"`
	Message          []Message `json:"message" gorm:"foreignkey:ChatRoomID"`
}

type Message struct {
	ID          int64     `json:"id" gorm:"primary_key"`
	ChatRoomID  int64     `json:"chat_room_id"`
	Content     string    `json:"content"`
	Attachment  string    `json:"attachement"`
	SenderID    int64     `json:"sender_id"`
	ReceiverID  int64     `json:"receiver_id"`
	SendAt      time.Time `json:"send_at"`
	SendAtIndex int64     `json:"-"`
}
