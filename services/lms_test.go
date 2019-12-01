package services

import (
	"reflect"
	"testing"

	"github.com/impal-lms/lms-backend/authentication"
	"github.com/impal-lms/lms-backend/repository"
)

func TestNewServices(t *testing.T) {
	type args struct {
		repo repository.Repository
		auth authentication.Authentication
	}
	tests := []struct {
		name string
		args args
		want *LMS
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServices(tt.args.repo, tt.args.auth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServices() = %v, want %v", got, tt.want)
			}
		})
	}
}
