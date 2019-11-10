package handler

import (
	"testing"

	"github.com/impal-lms/lms-backend/services"
	"github.com/labstack/echo/v4"
)

func TestHandler_GetAllUser(t *testing.T) {
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
			if err := h.GetAllUser(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetAllUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_CreateUser(t *testing.T) {
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
			if err := h.CreateUser(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_GetUserByID(t *testing.T) {
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
			if err := h.GetUserByID(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_UpdateUser(t *testing.T) {
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
			if err := h.UpdateUser(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_DeleteUserById(t *testing.T) {
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
			if err := h.DeleteUserById(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.DeleteUserById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_ChangePassword(t *testing.T) {
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
			if err := h.ChangePassword(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.ChangePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_ChangeRole(t *testing.T) {
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
			if err := h.ChangeRole(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.ChangeRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
