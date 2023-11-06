package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"problem1/models"
	"problem1/repository"
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
		{
			name: "success",
			fields: fields{
				repo: &repository.IUserRepositoryMock{
					FindFriendsByUserIDFunc: func(userID int) ([]*models.User, error) {
						return []*models.User{
							{ID: 2, Name: "User2", Friends: nil},
							{ID: 3, Name: "User3", Friends: nil},
						}, nil
					},
				},
			},
			args: args{userID: 1},
			want: []*models.User{
				{ID: 2, Name: "User2", Friends: nil},
				{ID: 3, Name: "User3", Friends: nil},
			},
			wantErr: false,
		},
		{
			name: "no user",
			fields: fields{
				repo: &repository.IUserRepositoryMock{
					FindFriendsByUserIDFunc: func(userID int) ([]*models.User, error) {
						return nil, ErrUserNotFound
					},
				},
			},
			args:    args{userID: 1},
			want:    nil,
			wantErr: true,
		},
		{
			name: "no friends",
			fields: fields{
				repo: &repository.IUserRepositoryMock{
					FindFriendsByUserIDFunc: func(userID int) ([]*models.User, error) {
						return []*models.User{}, nil
					},
				},
			},
			args:    args{userID: 1},
			want:    []*models.User{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				repo: tt.fields.repo,
			}
			got, err := u.GetFriendList(tt.args.userID)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
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
		{
			name: "success",
			fields: fields{
				repo: &repository.IUserRepositoryMock{
					FindFriendsOfFriendsByUserIDFunc: func(userID int) ([]*models.User, error) {
						return []*models.User{
							{ID: 2, Name: "User2", Friends: nil},
							{ID: 3, Name: "User3", Friends: nil},
						}, nil
					},
				},
			},
			args: args{userID: 1},
			want: []*models.User{
				{ID: 2, Name: "User2", Friends: nil},
				{ID: 3, Name: "User3", Friends: nil},
			},
			wantErr: false,
		},
		{
			name: "no user",
			fields: fields{
				repo: &repository.IUserRepositoryMock{
					FindFriendsOfFriendsByUserIDFunc: func(userID int) ([]*models.User, error) {
						return nil, ErrUserNotFound
					},
				},
			},
			args:    args{userID: 1},
			want:    nil,
			wantErr: true,
		},
		{
			name: "no friends",
			fields: fields{
				repo: &repository.IUserRepositoryMock{
					FindFriendsOfFriendsByUserIDFunc: func(userID int) ([]*models.User, error) {
						return []*models.User{}, nil
					},
				},
			},
			args:    args{userID: 1},
			want:    []*models.User{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				repo: tt.fields.repo,
			}
			got, err := u.GetFriendOfFriendList(tt.args.userID)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUserService_GetFriendsOfFriendsPagingByUserID(t *testing.T) {
	generateUsers := func(n int) []*models.User {
		var users []*models.User
		for i := 0; i < n; i++ {
			users = append(users, &models.User{
				ID:      uint64(i),
				Name:    fmt.Sprintf("User%d", i),
				Friends: nil,
			})
		}
		return users
	}
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
		{
			name: "success",
			fields: fields{
				repo: &repository.IUserRepositoryMock{
					FindFriendsOfFriendsPagingByUserIDFunc: func(userID int, page int, limit int) ([]*models.User, error) {
						return generateUsers(limit), nil
					},
				},
			},
			args:    args{userID: 1, page: 1, limit: 5},
			want:    generateUsers(5),
			wantErr: false,
		},
		{
			name: "no user",
			fields: fields{
				repo: &repository.IUserRepositoryMock{
					FindFriendsOfFriendsPagingByUserIDFunc: func(userID int, page int, limit int) ([]*models.User, error) {
						return nil, ErrUserNotFound
					},
				},
			},
			args:    args{userID: 1, page: 1, limit: 5},
			want:    nil,
			wantErr: true,
		},
		{
			name: "no friends",
			fields: fields{
				repo: &repository.IUserRepositoryMock{
					FindFriendsOfFriendsPagingByUserIDFunc: func(userID int, page int, limit int) ([]*models.User, error) {
						return []*models.User{}, nil
					},
				},
			},
			args:    args{},
			want:    []*models.User{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				repo: tt.fields.repo,
			}
			got, err := u.GetFriendOfFriendListPaging(tt.args.userID, tt.args.page, tt.args.limit)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
