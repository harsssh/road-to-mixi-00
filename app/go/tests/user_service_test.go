package tests

import (
	"problem1/models"
	"problem1/services"
	"reflect"
	"testing"
)

func TestUserService_GetFriendList(t *testing.T) {
	user1 := &models.User{ID: 1, UserID: 1, Name: "user1"}
	user2 := &models.User{ID: 2, UserID: 2, Name: "user2"}
	user3 := &models.User{ID: 3, UserID: 3, Name: "user3"}
	user1.Friends = []*models.User{user2}
	user2.Friends = []*models.User{user1}
	user3.Friends = []*models.User{}

	repo := NewMockUserRepository(map[int]*models.User{
		1: user1, 2: user2, 3: user3,
	})

	tests := []struct {
		name    string
		uid     int
		want    []*models.User
		wantErr bool
	}{
		{
			name:    "success",
			uid:     1,
			want:    []*models.User{user2},
			wantErr: false,
		},
		{
			name:    "no user",
			uid:     4,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "no friend",
			uid:     3,
			want:    []*models.User{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := services.NewUserService(repo)
			got, err := s.GetFriendList(tt.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFriendList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFriendList() got = %v, want %v", got, tt.want)
			}
		})
	}
}
