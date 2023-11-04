package services

import (
	"problem1/models"
	"reflect"
	"testing"
)

func TestUserService_GetFriendsByUserID(t *testing.T) {
	type fields struct {
		repo IUserRepository
	}
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				repo: tt.fields.repo,
			}
			got, err := u.GetFriendsByUserID(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFriendsByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFriendsByUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_GetFriendsOfFriendsByUserID(t *testing.T) {
	type fields struct {
		repo IUserRepository
	}
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				repo: tt.fields.repo,
			}
			got, err := u.GetFriendsOfFriendsByUserID(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFriendsOfFriendsByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFriendsOfFriendsByUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_GetFriendsOfFriendsPagingByUserID(t *testing.T) {
	type fields struct {
		repo IUserRepository
	}
	type args struct {
		userID int
		page   int
		limit  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				repo: tt.fields.repo,
			}
			got, err := u.GetFriendsOfFriendsPagingByUserID(tt.args.userID, tt.args.page, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFriendsOfFriendsPagingByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFriendsOfFriendsPagingByUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
