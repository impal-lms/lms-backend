package repository

import (
	"reflect"
	"testing"

	"github.com/impal-lms/lms-backend/domain"
	"github.com/jinzhu/gorm"
)

func TestGORM_GetAllChatRoom(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		userID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.ChatRoom
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.GetAllChatRoom(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.GetAllChatRoom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.GetAllChatRoom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_CreateChatRoom(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		ChatRoom domain.ChatRoom
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.ChatRoom
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.CreateChatRoom(tt.args.ChatRoom)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.CreateChatRoom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.CreateChatRoom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_GetChatRoomByID(t *testing.T) {
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
		want    domain.ChatRoom
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.GetChatRoomByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.GetChatRoomByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.GetChatRoomByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_UpdateChatRoom(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		ChatRoom domain.ChatRoom
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
			if err := g.UpdateChatRoom(tt.args.ChatRoom); (err != nil) != tt.wantErr {
				t.Errorf("GORM.UpdateChatRoom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGORM_DeleteChatRoomByID(t *testing.T) {
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
			if err := g.DeleteChatRoomByID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("GORM.DeleteChatRoomByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGORM_GetChatRoomBetweenUsers(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		studentID int64
		teacherID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.ChatRoom
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.GetChatRoomBetweenUsers(tt.args.studentID, tt.args.teacherID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.GetChatRoomBetweenUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.GetChatRoomBetweenUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_GetAllMessage(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		ChatRoomID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Message
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.GetAllMessage(tt.args.ChatRoomID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.GetAllMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.GetAllMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_CreateMessage(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		Message domain.Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Message
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.CreateMessage(tt.args.Message)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.CreateMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.CreateMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_GetMessageByID(t *testing.T) {
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
		want    domain.Message
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.GetMessageByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.GetMessageByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.GetMessageByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_UpdateMessage(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		Message domain.Message
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
			if err := g.UpdateMessage(tt.args.Message); (err != nil) != tt.wantErr {
				t.Errorf("GORM.UpdateMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGORM_DeleteMessageByID(t *testing.T) {
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
			if err := g.DeleteMessageByID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("GORM.DeleteMessageByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
