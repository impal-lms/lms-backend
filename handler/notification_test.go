package handler

import (
	"testing"

	"github.com/impal-lms/lms-backend/services"
	"github.com/labstack/echo/v4"
)

func TestHandler_GetAllNotification(t *testing.T) {
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
			if err := h.GetAllNotification(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetAllNotification() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_CreateNotification(t *testing.T) {
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
			if err := h.CreateNotification(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.CreateNotification() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_GetNotificationByID(t *testing.T) {
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
			if err := h.GetNotificationByID(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetNotificationByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_UpdateNotification(t *testing.T) {
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
			if err := h.UpdateNotification(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.UpdateNotification() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_DeleteNotificationByID(t *testing.T) {
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
			if err := h.DeleteNotificationByID(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.DeleteNotificationByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
