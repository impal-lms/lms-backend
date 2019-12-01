package handler

import (
	"testing"

	"github.com/impal-lms/lms-backend/services"
	"github.com/labstack/echo/v4"
)

func TestHandler_GetAllPresence(t *testing.T) {
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
			if err := h.GetAllPresence(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetAllPresence() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_CreatePresence(t *testing.T) {
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
			if err := h.CreatePresence(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.CreatePresence() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_GetPresenceByID(t *testing.T) {
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
			if err := h.GetPresenceByID(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetPresenceByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_UpdatePresence(t *testing.T) {
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
			if err := h.UpdatePresence(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.UpdatePresence() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_DeletePresenceByID(t *testing.T) {
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
			if err := h.DeletePresenceByID(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.DeletePresenceByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
