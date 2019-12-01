package domain

import "time"

type Presence struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	RFID      string    `json:"rfid"`
	StudentID int64     `json:"student_id"`
	Date      time.Time `json:"date"`
}
