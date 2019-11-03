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
