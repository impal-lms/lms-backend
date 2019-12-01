package repository

import (
	"reflect"
	"testing"

	"github.com/impal-lms/lms-backend/domain"
	"github.com/jinzhu/gorm"
)

func TestGORM_GetAllPresence(t *testing.T) {
	type fields struct {
		DB *gorm.DB
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
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.GetAllPresence(tt.args.studentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.GetAllPresence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.GetAllPresence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_CreatePresence(t *testing.T) {
	type fields struct {
		DB *gorm.DB
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
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.CreatePresence(tt.args.Presence)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.CreatePresence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.CreatePresence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_GetPresenceByID(t *testing.T) {
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
		want    domain.Presence
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.GetPresenceByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.GetPresenceByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.GetPresenceByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_UpdatePresence(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		Presence domain.Presence
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
			if err := g.UpdatePresence(tt.args.Presence); (err != nil) != tt.wantErr {
				t.Errorf("GORM.UpdatePresence() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGORM_DeletePresenceByID(t *testing.T) {
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
			if err := g.DeletePresenceByID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("GORM.DeletePresenceByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
