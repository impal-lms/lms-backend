package domain

type Classroom struct {
	ID         int64
	Label      string
	TeacherID  int64
	StudentIDs []int64
	RoomIDs    []int64
}
