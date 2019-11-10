package services

import (
	"reflect"
	"testing"

	"github.com/impal-lms/lms-backend/authentication"
	"github.com/impal-lms/lms-backend/domain"
	"github.com/impal-lms/lms-backend/repository"
)

func TestLMS_Authenticate(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.User
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
			got, err := lms.Authenticate(tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.Authenticate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_Login(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		req domain.LoginRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.LoginResponse
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
			got, err := lms.Login(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
