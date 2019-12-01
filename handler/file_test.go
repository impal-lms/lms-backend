package handler

import (
	"testing"

	"github.com/impal-lms/lms-backend/services"
	"github.com/labstack/echo/v4"
)

func TestHandler_GetAllMaterial(t *testing.T) {
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
			if err := h.GetAllMaterial(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetAllMaterial() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_CreateMaterial(t *testing.T) {
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
			if err := h.CreateMaterial(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.CreateMaterial() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_GetMaterialByID(t *testing.T) {
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
			if err := h.GetMaterialByID(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetMaterialByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_UpdateMaterial(t *testing.T) {
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
			if err := h.UpdateMaterial(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.UpdateMaterial() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_DeleteMaterialByID(t *testing.T) {
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
			if err := h.DeleteMaterialByID(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.DeleteMaterialByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_GetAllTask(t *testing.T) {
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
			if err := h.GetAllTask(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetAllTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_CreateTask(t *testing.T) {
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
			if err := h.CreateTask(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.CreateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_GetTaskByID(t *testing.T) {
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
			if err := h.GetTaskByID(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetTaskByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_UpdateTask(t *testing.T) {
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
			if err := h.UpdateTask(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_DeleteTaskByID(t *testing.T) {
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
			if err := h.DeleteTaskByID(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.DeleteTaskByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_GetAllSubmission(t *testing.T) {
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
			if err := h.GetAllSubmission(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetAllSubmission() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_CreateSubmission(t *testing.T) {
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
			if err := h.CreateSubmission(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.CreateSubmission() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_GetSubmissionByID(t *testing.T) {
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
			if err := h.GetSubmissionByID(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetSubmissionByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_UpdateSubmission(t *testing.T) {
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
			if err := h.UpdateSubmission(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.UpdateSubmission() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_DeleteSubmissionByID(t *testing.T) {
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
			if err := h.DeleteSubmissionByID(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.DeleteSubmissionByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
