package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/impal-lms/lms-backend/authentication"
	"github.com/impal-lms/lms-backend/domain"
	"github.com/impal-lms/lms-backend/repository"
	"github.com/impal-lms/lms-backend/services"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func TestIsValidExtension(t *testing.T) {
	type args struct {
		ext string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"pdf", args{ext: ".pdf"}, true},
		{"pdf.pdf", args{ext: ".pdf.pdf"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := services.IsValidExtension(tt.args.ext); got != tt.want {
				t.Errorf("IsValidExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_GetAllMaterial(t *testing.T) {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=127.0.0.1 port=5432 user=postgres dbname=lms password=postgres sslmode=disable"))
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	lms := services.NewServices(repository.NewRepository(db), authentication.NewAuthentication())
	classroomID0, err := lms.GetAllMaterial(0)

	type args struct {
		classroomID int64
	}

	tests := []struct {
		name    string
		args    args
		want    []domain.Material
		wantErr error
	}{
		{"ClassroomID: 0", args{classroomID: 0}, classroomID0, err},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lms.GetAllMaterial(tt.args.classroomID)
			if err != tt.wantErr {
				t.Errorf("LMS.GetAllMaterial() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetAllMaterial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_GetMaterialByID(t *testing.T) {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=127.0.0.1 port=5432 user=postgres dbname=lms password=postgres sslmode=disable"))
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	lms := services.NewServices(repository.NewRepository(db), authentication.NewAuthentication())
	MaterialID0, err := lms.GetMaterialByID(0)

	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		args    args
		want    domain.Material
		wantErr error
	}{
		{"MaterialID: 0", args{id: 0}, MaterialID0, err},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lms.GetMaterialByID(tt.args.id)
			if err != tt.wantErr {
				t.Errorf("LMS.GetMaterialByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetMaterialByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_GetAllTask(t *testing.T) {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=127.0.0.1 port=5432 user=postgres dbname=lms password=postgres sslmode=disable"))
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	lms := services.NewServices(repository.NewRepository(db), authentication.NewAuthentication())
	classroomID0, err := lms.GetAllTask(0)

	type args struct {
		classroomID int64
	}
	tests := []struct {
		name    string
		args    args
		want    []domain.Task
		wantErr error
	}{
		{"ClassroomID: 0", args{classroomID: 0}, classroomID0, err},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lms.GetAllTask(tt.args.classroomID)
			if err != tt.wantErr {
				t.Errorf("LMS.GetAllTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetAllTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_GetTaskByID(t *testing.T) {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=127.0.0.1 port=5432 user=postgres dbname=lms password=postgres sslmode=disable"))
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	lms := services.NewServices(repository.NewRepository(db), authentication.NewAuthentication())
	TaskID0, err := lms.GetTaskByID(0)

	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		args    args
		want    domain.Task
		wantErr error
	}{
		{"TaskID: 0", args{id: 0}, TaskID0, err},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lms.GetTaskByID(tt.args.id)
			if err != tt.wantErr {
				t.Errorf("LMS.GetTaskByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetTaskByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_GetAllSubmission(t *testing.T) {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=127.0.0.1 port=5432 user=postgres dbname=lms password=postgres sslmode=disable"))
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	lms := services.NewServices(repository.NewRepository(db), authentication.NewAuthentication())
	studentID0AndClassroomID0, err := lms.GetAllSubmission(0, 0)

	type args struct {
		studentID   int64
		classroomID int64
	}
	tests := []struct {
		name    string
		args    args
		want    []domain.Submission
		wantErr error
	}{
		{"StudentID: 0, ClassroomID: 0", args{classroomID: 0}, studentID0AndClassroomID0, err},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lms.GetAllSubmission(tt.args.studentID, tt.args.classroomID)
			if err != tt.wantErr {
				t.Errorf("LMS.GetAllSubmission() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetAllSubmission() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_GetSubmissionByID(t *testing.T) {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=127.0.0.1 port=5432 user=postgres dbname=lms password=postgres sslmode=disable"))
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	lms := services.NewServices(repository.NewRepository(db), authentication.NewAuthentication())
	SubmissionID0, err := lms.GetSubmissionByID(0)

	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		args    args
		want    domain.Submission
		wantErr error
	}{
		{"SubmissionID: 0", args{id: 0}, SubmissionID0, err},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lms.GetSubmissionByID(tt.args.id)
			if err != tt.wantErr {
				t.Errorf("LMS.GetSubmissionByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetSubmissionByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
