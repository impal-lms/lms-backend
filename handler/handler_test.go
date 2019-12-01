package handler

import (
	"reflect"
	"testing"

	"github.com/impal-lms/lms-backend/services"
	"github.com/labstack/echo/v4"
)

func TestNewHandler(t *testing.T) {
	type args struct {
		services services.Services
	}
	tests := []struct {
		name string
		args args
		want *Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandler(tt.args.services); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_HelloWorld(t *testing.T) {
	type fields struct {
		Services services.Services
	}
	type args struct {
		ctx echo.Context
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
			h := &Handler{
				Services: tt.fields.Services,
			}
			if err := h.HelloWorld(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.HelloWorld() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
