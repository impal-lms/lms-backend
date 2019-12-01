package services

import (
	"reflect"
	"testing"

	"github.com/impal-lms/lms-backend/authentication"
	"github.com/impal-lms/lms-backend/domain"
	"github.com/impal-lms/lms-backend/repository"
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidExtension(tt.args.ext); got != tt.want {
				t.Errorf("IsValidExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_GetAllMaterial(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		classroomID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Material
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
			got, err := lms.GetAllMaterial(tt.args.classroomID)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.GetAllMaterial() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetAllMaterial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_CreateMaterial(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		Material domain.Material
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Material
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
			got, got1, err := lms.CreateMaterial(tt.args.Material)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.CreateMaterial() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.CreateMaterial() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LMS.CreateMaterial() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLMS_GetMaterialByID(t *testing.T) {
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
		want    domain.Material
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
			got, err := lms.GetMaterialByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.GetMaterialByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetMaterialByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_UpdateMaterial(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		Material domain.Material
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Material
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
			got, err := lms.UpdateMaterial(tt.args.Material)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.UpdateMaterial() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.UpdateMaterial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_DeleteMaterialByID(t *testing.T) {
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
		want    domain.Material
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
			got, err := lms.DeleteMaterialByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.DeleteMaterialByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.DeleteMaterialByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_GetAllTask(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		classroomID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Task
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
			got, err := lms.GetAllTask(tt.args.classroomID)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.GetAllTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetAllTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_CreateTask(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		Task domain.Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Task
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
			got, got1, err := lms.CreateTask(tt.args.Task)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.CreateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.CreateTask() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LMS.CreateTask() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLMS_GetTaskByID(t *testing.T) {
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
		want    domain.Task
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
			got, err := lms.GetTaskByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.GetTaskByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetTaskByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_UpdateTask(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		Task domain.Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Task
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
			got, err := lms.UpdateTask(tt.args.Task)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.UpdateTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_DeleteTaskByID(t *testing.T) {
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
		want    domain.Task
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
			got, err := lms.DeleteTaskByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.DeleteTaskByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.DeleteTaskByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_GetAllSubmission(t *testing.T) {
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
		want    []domain.Submission
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
			got, err := lms.GetAllSubmission(tt.args.studentID, tt.args.classroomID)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.GetAllSubmission() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetAllSubmission() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_CreateSubmission(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		Submission domain.Submission
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Submission
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
			got, got1, err := lms.CreateSubmission(tt.args.Submission)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.CreateSubmission() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.CreateSubmission() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LMS.CreateSubmission() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLMS_GetSubmissionByID(t *testing.T) {
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
		want    domain.Submission
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
			got, err := lms.GetSubmissionByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.GetSubmissionByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetSubmissionByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_UpdateSubmission(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		Submission domain.Submission
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Submission
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
			got, err := lms.UpdateSubmission(tt.args.Submission)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.UpdateSubmission() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.UpdateSubmission() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_DeleteSubmissionByID(t *testing.T) {
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
		want    domain.Submission
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
			got, err := lms.DeleteSubmissionByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.DeleteSubmissionByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.DeleteSubmissionByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
