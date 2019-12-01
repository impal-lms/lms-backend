package services

import (
	"reflect"
	"testing"

	"github.com/impal-lms/lms-backend/authentication"
	"github.com/impal-lms/lms-backend/domain"
	"github.com/impal-lms/lms-backend/repository"
)

func TestLMS_GetAllChatRoom(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
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
			lms := &LMS{
				Repository:     tt.fields.Repository,
				Authentication: tt.fields.Authentication,
			}
			got, err := lms.GetAllChatRoom(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.GetAllChatRoom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetAllChatRoom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_GetChatRoomByID(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		ChatRoomID int64
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
			lms := &LMS{
				Repository:     tt.fields.Repository,
				Authentication: tt.fields.Authentication,
			}
			got, err := lms.GetChatRoomByID(tt.args.ChatRoomID)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.GetChatRoomByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetChatRoomByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_CreateChatRoom(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		ChatRoom domain.ChatRoom
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.ChatRoom
		want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lms := &LMS{
				Repository:     tt.fields.Repository,
				Authentication: tt.fields.Authentication,
			}
			got, got1, err := lms.CreateChatRoom(tt.args.ChatRoom)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.CreateChatRoom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.CreateChatRoom() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LMS.CreateChatRoom() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLMS_UpdateChatRoom(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
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
			lms := &LMS{
				Repository:     tt.fields.Repository,
				Authentication: tt.fields.Authentication,
			}
			got, err := lms.UpdateChatRoom(tt.args.ChatRoom)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.UpdateChatRoom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.UpdateChatRoom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_DeleteChatRoomByID(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		ChatRoomID int64
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
			lms := &LMS{
				Repository:     tt.fields.Repository,
				Authentication: tt.fields.Authentication,
			}
			got, err := lms.DeleteChatRoomByID(tt.args.ChatRoomID)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.DeleteChatRoomByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.DeleteChatRoomByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_GetAllMessage(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
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
			lms := &LMS{
				Repository:     tt.fields.Repository,
				Authentication: tt.fields.Authentication,
			}
			got, err := lms.GetAllMessage(tt.args.ChatRoomID)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.GetAllMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetAllMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_GetMessageByID(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		messageID int64
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
			lms := &LMS{
				Repository:     tt.fields.Repository,
				Authentication: tt.fields.Authentication,
			}
			got, err := lms.GetMessageByID(tt.args.messageID)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.GetMessageByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.GetMessageByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_CreateMessage(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		message domain.Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Message
		want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lms := &LMS{
				Repository:     tt.fields.Repository,
				Authentication: tt.fields.Authentication,
			}
			got, got1, err := lms.CreateMessage(tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.CreateMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.CreateMessage() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LMS.CreateMessage() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLMS_UpdateMessage(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		message domain.Message
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
			lms := &LMS{
				Repository:     tt.fields.Repository,
				Authentication: tt.fields.Authentication,
			}
			got, err := lms.UpdateMessage(tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.UpdateMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.UpdateMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMS_DeleteMessageByID(t *testing.T) {
	type fields struct {
		Repository     repository.Repository
		Authentication authentication.Authentication
	}
	type args struct {
		messageID int64
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
			lms := &LMS{
				Repository:     tt.fields.Repository,
				Authentication: tt.fields.Authentication,
			}
			got, err := lms.DeleteMessageByID(tt.args.messageID)
			if (err != nil) != tt.wantErr {
				t.Errorf("LMS.DeleteMessageByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMS.DeleteMessageByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
