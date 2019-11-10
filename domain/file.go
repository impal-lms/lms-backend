package domain

type Material struct {
	ID          int64  `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	FileURL     string `json:"file_url"`
	ClassroomID int64  `json:"classroom_id"`
}

type Task struct {
	ID          int64  `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	FileURL     string `json:"file_url"`
	ClassroomID int64  `json:"classroom_id"`
}

type Submission struct {
	ID        int64  `json:"id" gorm:"primary_key"`
	Title     string `json:"title"`
	FileURL   string `json:"file_url"`
	TaskID    int64  `json:"task_id"`
	StudentID int64  `json:"student_id"`
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

func IsValidExtension(ext string) bool {
	for _, all := range AllowedExtension {
		if ext == all {
			return true
		}
	}
	return false
}
