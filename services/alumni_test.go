package services

import (
	"reflect"
	"testing"

	"github.com/impal-lms/lms-backend/authentication"
	"github.com/impal-lms/lms-backend/domain"
	"github.com/impal-lms/lms-backend/repository"
)

func TestLMS_GetAllAlumni(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	tests := []struct {
		name    string
		fields  fields
		want    []domain.Alumni
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
			got, err := lms.GetAllAlumni()
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.GetAllAlumni() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetAllAlumni() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_CreateAlumni(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		Alumni domain.Alumni
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Alumni
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
			got, got1, err := lms.CreateAlumni(tt.args.Alumni)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.CreateAlumni() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.CreateAlumni() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LMS.CreateAlumni() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLMS_GetAlumniByID(t *testing.T) {
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
		want    domain.Alumni
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
			got, err := lms.GetAlumniByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.GetAlumniByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetAlumniByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_UpdateAlumni(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		Alumni domain.Alumni
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Alumni
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
			got, err := lms.UpdateAlumni(tt.args.Alumni)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.UpdateAlumni() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.UpdateAlumni() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_DeleteAlumniByID(t *testing.T) {
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
		want    domain.Alumni
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
			got, err := lms.DeleteAlumniByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.DeleteAlumniByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.DeleteAlumniByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
