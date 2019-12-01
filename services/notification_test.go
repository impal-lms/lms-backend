package services

import (
	"reflect"
	"testing"

	"github.com/impal-lms/lms-backend/authentication"
	"github.com/impal-lms/lms-backend/domain"
	"github.com/impal-lms/lms-backend/repository"
)

func TestLMS_GetAllNotification(t *testing.T) {
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
		want    []domain.Notification
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
			got, err := lms.GetAllNotification(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.GetAllNotification() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetAllNotification() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_CreateNotification(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		notification domain.Notification
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Notification
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
			got, got1, err := lms.CreateNotification(tt.args.notification)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.CreateNotification() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.CreateNotification() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LMS.CreateNotification() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLMS_GetNotificationByID(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		ID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Notification
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
			got, err := lms.GetNotificationByID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.GetNotificationByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetNotificationByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_UpdateNotification(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		notification domain.Notification
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Notification
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
			got, err := lms.UpdateNotification(tt.args.notification)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.UpdateNotification() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.UpdateNotification() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_DeleteNotificationByID(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		ID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Notification
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
			got, err := lms.DeleteNotificationByID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.DeleteNotificationByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.DeleteNotificationByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
