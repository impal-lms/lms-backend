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

func TestLMS_DeleteClassroomById(t *testing.T) {
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
			got, err := lms.DeleteClassroomById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.DeleteClassroomById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.DeleteClassroomById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_AddStudentToClassroom(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		id        int64
		studentID int64
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
			got, err := lms.AddStudentToClassroom(tt.args.id, tt.args.studentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.AddStudentToClassroom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.AddStudentToClassroom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_DeleteStudentFromClassroom(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		id        int64
		studentID int64
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
			got, err := lms.DeleteStudentFromClassroom(tt.args.id, tt.args.studentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.DeleteStudentFromClassroom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.DeleteStudentFromClassroom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_AddRoomToClassroom(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		id     int64
		roomID int64
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
			got, err := lms.AddRoomToClassroom(tt.args.id, tt.args.roomID)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.AddRoomToClassroom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.AddRoomToClassroom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_DeleteRoomFromClassroom(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		id     int64
		roomID int64
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
			got, err := lms.DeleteRoomFromClassroom(tt.args.id, tt.args.roomID)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.DeleteRoomFromClassroom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.DeleteRoomFromClassroom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_GetAllClassroomOfUser(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		userID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.ClassroomOfUser
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
			got, err := lms.GetAllClassroomOfUser(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.GetAllClassroomOfUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetAllClassroomOfUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
