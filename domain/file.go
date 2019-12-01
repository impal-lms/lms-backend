package domain

import "time"

type Material struct {
	ID          int64  `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	FileURL     string `json:"file_url"`
	ClassroomID int64  `json:"classroom_id"`
}

type Task struct {
	ID          int64     `json:"id" gorm:"primary_key"`
	Title       string    `json:"title"`
	FileURL     string    `json:"file_url"`
	DueDate     time.Time `json:"due_date"`
	ClassroomID int64     `json:"classroom_id"`
}

type Submission struct {
	ID         int64     `json:"id" gorm:"primary_key"`
	Title      string    `json:"title"`
	FileURL    string    `json:"file_url"`
	SubmitDate time.Time `json:"submitDate"`
	TaskID     int64     `json:"task_id"`
	StudentID  int64     `json:"student_id"`
}

var AllowedExtension = []string{
	".pdf",
	".doc",
	".docx",
	".xls",
	".xlsx",
	".ppt",
	".pptx",
}
