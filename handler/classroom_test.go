package handler

import (
	"testing"

	"github.com/impal-lms/lms-backend/services"
	"github.com/labstack/echo/v4"
)

func TestHandler_GetAllClassroom(t *testing.T) {
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
			if err := h.GetAllClassroom(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetAllClassroom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_CreateClassroom(t *testing.T) {
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
			if err := h.CreateClassroom(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.CreateClassroom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_GetClassroomByID(t *testing.T) {
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
			if err := h.GetClassroomByID(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetClassroomByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_UpdateClassroom(t *testing.T) {
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
			if err := h.UpdateClassroom(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.UpdateClassroom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_DeleteClassroomById(t *testing.T) {
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
			if err := h.DeleteClassroomById(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.DeleteClassroomById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_AddStudentToClassroom(t *testing.T) {
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
			if err := h.AddStudentToClassroom(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.AddStudentToClassroom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_DeleteStudentFromClassroom(t *testing.T) {
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
			if err := h.DeleteStudentFromClassroom(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.DeleteStudentFromClassroom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_AddRoomToClassroom(t *testing.T) {
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
			if err := h.AddRoomToClassroom(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.AddRoomToClassroom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_DeleteRoomFromClassroom(t *testing.T) {
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
			if err := h.DeleteRoomFromClassroom(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.DeleteRoomFromClassroom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_GetAllClassroomOfUser(t *testing.T) {
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
			if err := h.GetAllClassroomOfUser(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetAllClassroomOfUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
