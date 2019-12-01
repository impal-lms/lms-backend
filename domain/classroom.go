package domain

type Classroom struct {
	ID               int64              `json:"id" gorm:"primary_key"`
	Label            string             `json:"label"`
	Subject          string             `json:"subject"`
	Room             string             `json:"room"`
	DaySchedule      int                `json:"day_schedule"`
	TeacherID        int64              `json:"teacher_id"`
	StudentClassroom []StudentClassroom `json:"student_classroom" gorm:"foreignkey:ClassroomID"`
}

type StudentClassroom struct {
	ID          int64 `json:"id" gorm:"primary_key"`
	StudentID   int64 `json:"student_id"`
	ClassroomID int64 `json:"classroom_id"`
}
