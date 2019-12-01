package repository

import (
	"reflect"
	"testing"

	"github.com/impal-lms/lms-backend/domain"
	"github.com/jinzhu/gorm"
)

func TestGORM_GetAllAlumni(t *testing.T) {
	type fields struct {
		DB *gorm.DB
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
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.GetAllAlumni()
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.GetAllAlumni() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.GetAllAlumni() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_CreateAlumni(t *testing.T) {
	type fields struct {
		DB *gorm.DB
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
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.CreateAlumni(tt.args.Alumni)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.CreateAlumni() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.CreateAlumni() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_GetAlumniByID(t *testing.T) {
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
		want    domain.Alumni
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.GetAlumniByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.GetAlumniByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.GetAlumniByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_UpdateAlumni(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		Alumni domain.Alumni
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
			if err := g.UpdateAlumni(tt.args.Alumni); (err != nil) != tt.wantErr {
				t.Errorf("GORM.UpdateAlumni() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGORM_DeleteAlumniByID(t *testing.T) {
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
			if err := g.DeleteAlumniByID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("GORM.DeleteAlumniByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
