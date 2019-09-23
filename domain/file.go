package domain

type Material struct {
	ID          int64
	FileURL     string
	ClassroomID int64
}

type Task struct {
	ID          int64
	FileURL     string
	ClassRoomID int64
}

type Submission struct {
	ID        int64
	FileURL   string
	TaskID    int64
	StudentID int64
}
