package repository

import (
	"reflect"
	"testing"

	"github.com/impal-lms/lms-backend/domain"
	"github.com/jinzhu/gorm"
)

func TestGORM_GetAllNotification(t *testing.T) {
	type fields struct {
		DB *gorm.DB
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
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.GetAllNotification(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.GetAllNotification() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.GetAllNotification() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_CreateNotification(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		Notification domain.Notification
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
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.CreateNotification(tt.args.Notification)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.CreateNotification() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.CreateNotification() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_GetNotificationByID(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		id int64
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
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.GetNotificationByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.GetNotificationByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.GetNotificationByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_UpdateNotification(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		Notification domain.Notification
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			if err := g.UpdateNotification(tt.args.Notification); (err != nil) != tt.wantErr {
				t.Errorf("GORM.UpdateNotification() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGORM_DeleteNotificationByID(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			if err := g.DeleteNotificationByID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("GORM.DeleteNotificationByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
