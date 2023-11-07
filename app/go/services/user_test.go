package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"problem1/mocks"
	"problem1/models"
	"strconv"
	"testing"
)

const totalUser = 5

func getUsers(n int) []*models.User {
	users := make([]*models.User, n)
	for i := 0; i < n; i++ {
		users[i] = &models.User{ID: int64(i), UserID: i, Name: "user" + strconv.Itoa(i)}
	}
	return users
}

func TestUserService_GetFriendsByUserID(t *testing.T) {
	type fields struct {
		repo IUserRepository
	}
	type args struct {
		userID int
	}
	users := getUsers(totalUser)
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
				repo: &mocks.IUserRepositoryMock{
					FindFriendsByUserIDFunc: func(userID int) ([]*models.User, error) {
						return users, nil
					},
				},
			},
			args:    args{userID: 1},
			want:    users,
			wantErr: false,
		},
		{
			name: "no user",
			fields: fields{
				repo: &mocks.IUserRepositoryMock{
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
				repo: &mocks.IUserRepositoryMock{
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
	users := getUsers(totalUser)
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
				repo: &mocks.IUserRepositoryMock{
					FindFriendsOfFriendsByUserIDFunc: func(userID int) ([]*models.User, error) {
						return users, nil
					},
					FindFriendsByUserIDFunc: func(userID int) ([]*models.User, error) {
						return users[:1], nil
					},
					FindBlockUsersByUserIDFunc: func(userID int) ([]*models.User, error) {
						return users[1:2], nil
					},
				},
			},
			args:    args{userID: 1},
			want:    users[2:],
			wantErr: false,
		},
		{
			name: "no user",
			fields: fields{
				repo: &mocks.IUserRepositoryMock{
					FindFriendsOfFriendsByUserIDFunc: func(userID int) ([]*models.User, error) {
						return nil, ErrUserNotFound
					},
					FindFriendsByUserIDFunc: func(userID int) ([]*models.User, error) {
						return nil, nil
					},
					FindBlockUsersByUserIDFunc: func(userID int) ([]*models.User, error) {
						return nil, nil
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
				repo: &mocks.IUserRepositoryMock{
					FindFriendsOfFriendsByUserIDFunc: func(userID int) ([]*models.User, error) {
						return []*models.User{}, nil
					},
					FindFriendsByUserIDFunc: func(userID int) ([]*models.User, error) {
						return nil, nil
					},
					FindBlockUsersByUserIDFunc: func(userID int) ([]*models.User, error) {
						return nil, nil
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
	type fields struct {
		repo IUserRepository
	}
	type args struct {
		userID int
		page   int
		limit  int
	}
	users := getUsers(totalUser)
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        []*models.User
		wantErr     bool
		assertCalls func(mock *mocks.IUserRepositoryMock)
	}{
		{
			name: "success",
			fields: fields{
				repo: &mocks.IUserRepositoryMock{
					FindFriendsByUserIDFunc: func(userID int) ([]*models.User, error) {
						return users[:1], nil
					},
					FindBlockUsersByUserIDFunc: func(userID int) ([]*models.User, error) {
						return users[1:2], nil
					},
					FindFriendsOfFriendsExcludingSomeUsersByUserIDWithPaginationFunc: func(
						userID int, excludedUserIDs []int, page int, limit int,
					) ([]*models.User, error) {
						return users[2:], nil
					},
				},
			},
			args:    args{userID: 1, page: 1, limit: 5},
			want:    users[2:],
			wantErr: false,
			assertCalls: func(mock *mocks.IUserRepositoryMock) {
				calls := mock.FindFriendsOfFriendsExcludingSomeUsersByUserIDWithPaginationCalls()
				assert.Equal(t, 1, len(calls))
				assert.Equal(t, []int{0, 1}, calls[0].ExcludedUserIDs)
			},
		},
		{
			name: "no user",
			fields: fields{
				repo: &mocks.IUserRepositoryMock{
					FindFriendsByUserIDFunc: func(userID int) ([]*models.User, error) {
						return nil, nil
					},
					FindBlockUsersByUserIDFunc: func(userID int) ([]*models.User, error) {
						return nil, nil
					},
					FindFriendsOfFriendsExcludingSomeUsersByUserIDWithPaginationFunc: func(
						userID int, excludedUserIDs []int, page int, limit int,
					) ([]*models.User, error) {
						return nil, ErrUserNotFound
					},
				},
			},
			args:        args{userID: 1, page: 1, limit: 5},
			want:        nil,
			wantErr:     true,
			assertCalls: nil,
		},
		{
			name: "no friends",
			fields: fields{
				repo: &mocks.IUserRepositoryMock{
					FindFriendsByUserIDFunc: func(userID int) ([]*models.User, error) {
						return nil, nil
					},
					FindBlockUsersByUserIDFunc: func(userID int) ([]*models.User, error) {
						return nil, nil
					},
					FindFriendsOfFriendsExcludingSomeUsersByUserIDWithPaginationFunc: func(
						userID int, excludedUserIDs []int, page int, limit int,
					) ([]*models.User, error) {
						return []*models.User{}, nil
					},
				},
			},
			args:        args{userID: 1, page: 1, limit: 5},
			want:        []*models.User{},
			wantErr:     false,
			assertCalls: nil,
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
			if tt.assertCalls != nil {
				tt.assertCalls(tt.fields.repo.(*mocks.IUserRepositoryMock))
			}
		})
	}
}
