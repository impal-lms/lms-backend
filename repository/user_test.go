package repository

import (
	"reflect"
	"testing"

	"github.com/impal-lms/lms-backend/domain"
	"github.com/jinzhu/gorm"
)

func TestGORM_GetAllUser(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []domain.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.GetAllUser()
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.GetAllUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.GetAllUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_CreateUser(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		user domain.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.CreateUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_GetUserByID(t *testing.T) {
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
		want    domain.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.GetUserByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.GetUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_GetUserByEmail(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.GetUserByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.GetUserByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_UpdateUser(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		user domain.User
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
			if err := g.UpdateUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("GORM.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGORM_DeleteUserById(t *testing.T) {
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
			if err := g.DeleteUserById(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("GORM.DeleteUserById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
