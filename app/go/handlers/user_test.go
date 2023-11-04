package handlers

import (
	"github.com/labstack/echo/v4"
	"problem1/services"
	"testing"
)

func TestUserHandler_GetFriendList(t *testing.T) {
	type fields struct {
		service services.IUserService
	}
	type args struct {
		c echo.Context
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
			h := &UserHandler{
				service: tt.fields.service,
			}
			if err := h.GetFriendList(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetFriendList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserHandler_GetFriendOfFriendList(t *testing.T) {
	type fields struct {
		service services.IUserService
	}
	type args struct {
		c echo.Context
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
			h := &UserHandler{
				service: tt.fields.service,
			}
			if err := h.GetFriendOfFriendList(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetFriendOfFriendList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserHandler_GetFriendOfFriendListPaging(t *testing.T) {
	type fields struct {
		service services.IUserService
	}
	type args struct {
		c echo.Context
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
			h := &UserHandler{
				service: tt.fields.service,
			}
			if err := h.GetFriendOfFriendListPaging(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetFriendOfFriendListPaging() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
