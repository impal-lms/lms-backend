package domain

type Classroom struct {
	ID         int64   `json:"id" gorm:"primary_key"`
	Label      string  `json:"label"`
	TeacherID  int64   `json:"teacher_id"`
	StudentIDs []int64 `json:"student_ids"`
	RoomIDs    []int64 `json:"room_ids"`
}
