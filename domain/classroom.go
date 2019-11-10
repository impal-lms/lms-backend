package domain

import "github.com/lib/pq"

type Classroom struct {
	ID         int64         `json:"id" gorm:"primary_key"`
	Label      string        `json:"label"`
	TeacherID  int64         `json:"teacher_id"`
	StudentIDs pq.Int64Array `json:"student_ids" gorm:"type:BigInt[]"`
	RoomIDs    pq.Int64Array `json:"room_ids" gorm:"type:BigInt[]"`
}

type CreateUpdateClassroomRequest struct {
	Label     string `json:"label"`
	TeacherID int64  `json:"teacher_id"`
}

type AddStudentToClassroomRequest struct {
	StudentID int64 `json:"student_id"`
}

type AddRoomToClassroomRequest struct {
	RoomID int64 `json:"room_id"`
}

type ClassroomOfUser struct {
	ID        int64  `json:"id" gorm:"primary_key"`
	Label     string `json:"label"`
	TeacherID int64  `json:"teacher_id"`
}
