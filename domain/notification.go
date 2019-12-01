package domain

import "time"

type Notification struct {
	ID          int64     `json:"id" gorm:"primary_key"`
	SenderID    int64     `json:"sender_id"`
	ReceiverID  int64     `json:"receiver_id"`
	Content     string    `json:"content"`
	SendAt      time.Time `json:"send_at"`
	SendAtIndex int64     `json:"-"`
}
