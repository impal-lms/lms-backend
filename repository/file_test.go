package repository

import (
	"reflect"
	"testing"

	"github.com/impal-lms/lms-backend/domain"
	"github.com/jinzhu/gorm"
)

func TestGORM_GetAllMaterial(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		classroomID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Material
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.GetAllMaterial(tt.args.classroomID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.GetAllMaterial() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.GetAllMaterial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_CreateMaterial(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		Material domain.Material
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Material
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.CreateMaterial(tt.args.Material)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.CreateMaterial() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.CreateMaterial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_GetMaterialByID(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Material
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.GetMaterialByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.GetMaterialByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.GetMaterialByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_UpdateMaterial(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		Material domain.Material
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
			g := &GORM{
				DB: tt.fields.DB,
			}
			if err := g.UpdateMaterial(tt.args.Material); (err != nil) != tt.wantErr {
				t.Errorf("GORM.UpdateMaterial() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGORM_DeleteMaterialById(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		id int64
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
			g := &GORM{
				DB: tt.fields.DB,
			}
			if err := g.DeleteMaterialById(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("GORM.DeleteMaterialById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGORM_GetAllTask(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		classroomID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.GetAllTask(tt.args.classroomID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.GetAllTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.GetAllTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_CreateTask(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		Task domain.Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.CreateTask(tt.args.Task)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.CreateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.CreateTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_GetTaskByID(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.GetTaskByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.GetTaskByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.GetTaskByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_UpdateTask(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		Task domain.Task
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
			g := &GORM{
				DB: tt.fields.DB,
			}
			if err := g.UpdateTask(tt.args.Task); (err != nil) != tt.wantErr {
				t.Errorf("GORM.UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGORM_DeleteTaskById(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		id int64
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
			g := &GORM{
				DB: tt.fields.DB,
			}
			if err := g.DeleteTaskById(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("GORM.DeleteTaskById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGORM_GetAllSubmission(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		classroomID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Submission
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.GetAllSubmission(tt.args.classroomID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.GetAllSubmission() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.GetAllSubmission() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_CreateSubmission(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		Submission domain.Submission
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Submission
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.CreateSubmission(tt.args.Submission)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.CreateSubmission() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.CreateSubmission() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_GetSubmissionByID(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Submission
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.GetSubmissionByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.GetSubmissionByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.GetSubmissionByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_UpdateSubmission(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		Submission domain.Submission
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
			g := &GORM{
				DB: tt.fields.DB,
			}
			if err := g.UpdateSubmission(tt.args.Submission); (err != nil) != tt.wantErr {
				t.Errorf("GORM.UpdateSubmission() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGORM_DeleteSubmissionById(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		id int64
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
			g := &GORM{
				DB: tt.fields.DB,
			}
			if err := g.DeleteSubmissionById(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("GORM.DeleteSubmissionById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
