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

func TestUserService_GetFriendsByID(t *testing.T) {
	type fields struct {
		repo IUserRepository
	}
	type args struct {
		id int64
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
					FindFriendsByIDFunc: func(id int64) ([]*models.User, error) {
						return users, nil
					},
				},
			},
			args:    args{id: 1},
			want:    users,
			wantErr: false,
		},
		{
			name: "no user",
			fields: fields{
				repo: &mocks.IUserRepositoryMock{
					FindFriendsByIDFunc: func(id int64) ([]*models.User, error) {
						return nil, ErrUserNotFound
					},
				},
			},
			args:    args{id: 1},
			want:    nil,
			wantErr: true,
		},
		{
			name: "no friends",
			fields: fields{
				repo: &mocks.IUserRepositoryMock{
					FindFriendsByIDFunc: func(id int64) ([]*models.User, error) {
						return []*models.User{}, nil
					},
				},
			},
			args:    args{id: 1},
			want:    []*models.User{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				repo: tt.fields.repo,
			}
			got, err := u.GetFriendList(tt.args.id)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUserService_GetFriendsOfFriendsByID(t *testing.T) {
	type fields struct {
		repo IUserRepository
	}
	type args struct {
		id int64
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
					FindFriendsOfFriendsByIDFunc: func(id int64) ([]*models.User, error) {
						return users, nil
					},
					FindFriendsByIDFunc: func(id int64) ([]*models.User, error) {
						return users[:1], nil
					},
					FindBlockUsersByIDFunc: func(id int64) ([]*models.User, error) {
						return users[1:2], nil
					},
				},
			},
			args:    args{id: 1},
			want:    users[2:],
			wantErr: false,
		},
		{
			name: "no user",
			fields: fields{
				repo: &mocks.IUserRepositoryMock{
					FindFriendsOfFriendsByIDFunc: func(id int64) ([]*models.User, error) {
						return nil, ErrUserNotFound
					},
					FindFriendsByIDFunc: func(id int64) ([]*models.User, error) {
						return nil, nil
					},
					FindBlockUsersByIDFunc: func(id int64) ([]*models.User, error) {
						return nil, nil
					},
				},
			},
			args:    args{id: 1},
			want:    nil,
			wantErr: true,
		},
		{
			name: "no friends",
			fields: fields{
				repo: &mocks.IUserRepositoryMock{
					FindFriendsOfFriendsByIDFunc: func(id int64) ([]*models.User, error) {
						return []*models.User{}, nil
					},
					FindFriendsByIDFunc: func(id int64) ([]*models.User, error) {
						return nil, nil
					},
					FindBlockUsersByIDFunc: func(id int64) ([]*models.User, error) {
						return nil, nil
					},
				},
			},
			args:    args{id: 1},
			want:    []*models.User{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				repo: tt.fields.repo,
			}
			got, err := u.GetFriendOfFriendList(tt.args.id)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUserService_GetFriendsOfFriendsPagingByID(t *testing.T) {
	type fields struct {
		repo IUserRepository
	}
	type args struct {
		id    int64
		page  int
		limit int
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
					FindFriendsByIDFunc: func(id int64) ([]*models.User, error) {
						return users[:1], nil
					},
					FindBlockUsersByIDFunc: func(id int64) ([]*models.User, error) {
						return users[1:2], nil
					},
					FindFriendsOfFriendsExcludingSomeUsersByIDWithPaginationFunc: func(
						id int64, excludeIDs []int64, page int, limit int,
					) ([]*models.User, error) {
						return users[2:], nil
					},
				},
			},
			args:    args{id: 1, page: 1, limit: 5},
			want:    users[2:],
			wantErr: false,
			assertCalls: func(mock *mocks.IUserRepositoryMock) {
				calls := mock.FindFriendsOfFriendsExcludingSomeUsersByIDWithPaginationCalls()
				assert.Equal(t, 1, len(calls))
				assert.Equal(t, []int64{0, 1}, calls[0].ExcludeIDs)
			},
		},
		{
			name: "no user",
			fields: fields{
				repo: &mocks.IUserRepositoryMock{
					FindFriendsByIDFunc: func(id int64) ([]*models.User, error) {
						return nil, nil
					},
					FindBlockUsersByIDFunc: func(id int64) ([]*models.User, error) {
						return nil, nil
					},
					FindFriendsOfFriendsExcludingSomeUsersByIDWithPaginationFunc: func(
						id int64, excludeIDs []int64, page int, limit int,
					) ([]*models.User, error) {
						return nil, ErrUserNotFound
					},
				},
			},
			args:        args{id: 1, page: 1, limit: 5},
			want:        nil,
			wantErr:     true,
			assertCalls: nil,
		},
		{
			name: "no friends",
			fields: fields{
				repo: &mocks.IUserRepositoryMock{
					FindFriendsByIDFunc: func(id int64) ([]*models.User, error) {
						return nil, nil
					},
					FindBlockUsersByIDFunc: func(id int64) ([]*models.User, error) {
						return nil, nil
					},
					FindFriendsOfFriendsExcludingSomeUsersByIDWithPaginationFunc: func(
						id int64, excludeIDs []int64, page int, limit int,
					) ([]*models.User, error) {
						return []*models.User{}, nil
					},
				},
			},
			args:        args{id: 1, page: 1, limit: 5},
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
			got, err := u.GetFriendOfFriendListPaging(tt.args.id, tt.args.page, tt.args.limit)
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
