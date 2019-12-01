package services

import (
	"reflect"
	"testing"

	"github.com/impal-lms/lms-backend/authentication"
	"github.com/impal-lms/lms-backend/domain"
	"github.com/impal-lms/lms-backend/repository"
)

func TestLMS_GetAllClassroom(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	tests := []struct {
		name    string
		fields  fields
		want    []domain.Classroom
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lms := &LMS{
				Repository:     tt.fields.Repository,
				Authentication: tt.fields.Authentication,
			}
			got, err := lms.GetAllClassroom()
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.GetAllClassroom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetAllClassroom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_CreateClassroom(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		Classroom domain.Classroom
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Classroom
		want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lms := &LMS{
				Repository:     tt.fields.Repository,
				Authentication: tt.fields.Authentication,
			}
			got, got1, err := lms.CreateClassroom(tt.args.Classroom)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.CreateClassroom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.CreateClassroom() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LMS.CreateClassroom() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLMS_GetClassroomByID(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Classroom
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lms := &LMS{
				Repository:     tt.fields.Repository,
				Authentication: tt.fields.Authentication,
			}
			got, err := lms.GetClassroomByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.GetClassroomByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetClassroomByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_UpdateClassroom(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		Classroom domain.Classroom
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Classroom
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lms := &LMS{
				Repository:     tt.fields.Repository,
				Authentication: tt.fields.Authentication,
			}
			got, err := lms.UpdateClassroom(tt.args.Classroom)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.UpdateClassroom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.UpdateClassroom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_DeleteClassroomByID(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Classroom
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lms := &LMS{
				Repository:     tt.fields.Repository,
				Authentication: tt.fields.Authentication,
			}
			got, err := lms.DeleteClassroomByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.DeleteClassroomByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.DeleteClassroomByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_GetAllStudentClassroom(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		studentID   int64
		classroomID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.StudentClassroom
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lms := &LMS{
				Repository:     tt.fields.Repository,
				Authentication: tt.fields.Authentication,
			}
			got, err := lms.GetAllStudentClassroom(tt.args.studentID, tt.args.classroomID)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.GetAllStudentClassroom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetAllStudentClassroom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_CreateStudentClassroom(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		StudentClassroom domain.StudentClassroom
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.StudentClassroom
		want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lms := &LMS{
				Repository:     tt.fields.Repository,
				Authentication: tt.fields.Authentication,
			}
			got, got1, err := lms.CreateStudentClassroom(tt.args.StudentClassroom)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.CreateStudentClassroom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.CreateStudentClassroom() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LMS.CreateStudentClassroom() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLMS_GetStudentClassroomByID(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.StudentClassroom
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lms := &LMS{
				Repository:     tt.fields.Repository,
				Authentication: tt.fields.Authentication,
			}
			got, err := lms.GetStudentClassroomByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.GetStudentClassroomByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetStudentClassroomByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_UpdateStudentClassroom(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		StudentClassroom domain.StudentClassroom
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.StudentClassroom
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lms := &LMS{
				Repository:     tt.fields.Repository,
				Authentication: tt.fields.Authentication,
			}
			got, err := lms.UpdateStudentClassroom(tt.args.StudentClassroom)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.UpdateStudentClassroom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.UpdateStudentClassroom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_DeleteStudentClassroomByID(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.StudentClassroom
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lms := &LMS{
				Repository:     tt.fields.Repository,
				Authentication: tt.fields.Authentication,
			}
			got, err := lms.DeleteStudentClassroomByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.DeleteStudentClassroomByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.DeleteStudentClassroomByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
