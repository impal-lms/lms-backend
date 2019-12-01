package repository

import (
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
)

func TestNewRepository(t *testing.T) {
	type args struct {
		DB *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *GORM
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRepository(tt.args.DB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
