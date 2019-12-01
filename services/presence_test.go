package services

import (
	"reflect"
	"testing"

	"github.com/impal-lms/lms-backend/authentication"
	"github.com/impal-lms/lms-backend/domain"
	"github.com/impal-lms/lms-backend/repository"
)

func TestLMS_GetAllPresence(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		studentID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Presence
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
			got, err := lms.GetAllPresence(tt.args.studentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.GetAllPresence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetAllPresence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_CreatePresence(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		Presence domain.Presence
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Presence
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
			got, got1, err := lms.CreatePresence(tt.args.Presence)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.CreatePresence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.CreatePresence() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LMS.CreatePresence() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLMS_GetPresenceByID(t *testing.T) {
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
		want    domain.Presence
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
			got, err := lms.GetPresenceByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.GetPresenceByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetPresenceByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_UpdatePresence(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		Presence domain.Presence
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Presence
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
			got, err := lms.UpdatePresence(tt.args.Presence)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.UpdatePresence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.UpdatePresence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_DeletePresenceByID(t *testing.T) {
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
		want    domain.Presence
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
			got, err := lms.DeletePresenceByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.DeletePresenceByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.DeletePresenceByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
