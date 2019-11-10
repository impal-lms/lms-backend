package services

import (
	"reflect"
	"testing"

	"github.com/impal-lms/lms-backend/authentication"
	"github.com/impal-lms/lms-backend/domain"
	"github.com/impal-lms/lms-backend/repository"
)

func TestLMS_GetAllUser(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
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
			lms := &LMS{
				Repository:     tt.fields.Repository,
				Authentication: tt.fields.Authentication,
			}
			got, err := lms.GetAllUser()
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.GetAllUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetAllUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_CreateUser(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		user domain.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.User
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
			got, got1, err := lms.CreateUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.CreateUser() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LMS.CreateUser() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLMS_GetUserByID(t *testing.T) {
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
		want    domain.User
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
			got, err := lms.GetUserByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_GetUserByEmail(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
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
			lms := &LMS{
				Repository:     tt.fields.Repository,
				Authentication: tt.fields.Authentication,
			}
			got, err := lms.GetUserByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetUserByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_UpdateUser(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
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
			lms := &LMS{
				Repository:     tt.fields.Repository,
				Authentication: tt.fields.Authentication,
			}
			got, err := lms.UpdateUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_DeleteUserById(t *testing.T) {
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
		want    domain.User
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
			got, err := lms.DeleteUserById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.DeleteUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.DeleteUserById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_ChangePassword(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		user domain.User
		req  domain.ChangePasswordRequest
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
			lms := &LMS{
				Repository:     tt.fields.Repository,
				Authentication: tt.fields.Authentication,
			}
			got, err := lms.ChangePassword(tt.args.user, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.ChangePassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.ChangePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_ChangeRole(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		user domain.User
		req  domain.ChangeRoleRequest
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
			lms := &LMS{
				Repository:     tt.fields.Repository,
				Authentication: tt.fields.Authentication,
			}
			got, err := lms.ChangeRole(tt.args.user, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.ChangeRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.ChangeRole() = %v, want %v", got, tt.want)
			}
		})
	}
}
