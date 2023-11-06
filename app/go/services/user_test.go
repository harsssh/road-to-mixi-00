package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"problem1/mocks"
	"problem1/models"
	"strconv"
	"testing"
)

const TotalUsersOfMockRepository = 6

// NOTE: `FindFriendsOfFriendsByUserID` returns valid result only for user1
// user0 has no friends
// friend: (1, 2), (1, 3), (2, 3), (2, 4), (2, 5)
// blocked: (1, 5)
func getUserRepositoryMock() (IUserRepository, []*models.User) {
	users := make([]*models.User, TotalUsersOfMockRepository)
	for i := range users {
		users[i] = &models.User{ID: uint64(i), UserID: i, Name: "user" + strconv.Itoa(i)}
	}
	users[1].Friends = []*models.User{users[2], users[3]}
	users[1].Blocks = []*models.User{users[5]}
	users[2].Friends = []*models.User{users[1], users[3], users[4], users[5]}
	users[3].Friends = []*models.User{users[1], users[2]}
	users[4].Friends = []*models.User{users[2]}
	return &mocks.IUserRepositoryMock{
		FindFriendsByUserIDFunc: func(userID int) ([]*models.User, error) {
			if userID < 0 || userID >= TotalUsersOfMockRepository {
				return nil, ErrUserNotFound
			}
			if users[userID].Friends == nil {
				return []*models.User{}, nil
			}
			return users[userID].Friends, nil
		},
		FindBlockUsersByUserIDFunc: func(userID int) ([]*models.User, error) {
			if userID < 0 || userID >= TotalUsersOfMockRepository {
				return nil, ErrUserNotFound
			}
			if users[userID].Blocks == nil {
				return []*models.User{}, nil
			}
			return users[userID].Blocks, nil
		},
		FindFriendsOfFriendsByUserIDFunc: func(userID int) ([]*models.User, error) {
			if userID == 0 {
				return []*models.User{}, nil
			}
			return []*models.User{users[2], users[3], users[4], users[5]}, nil
		},
		FindFriendsOfFriendsExcludingSomeUsersByUserIDWithPaginationFunc: func(
			userID int, excludedUserIDs []int, page int, limit int,
		) ([]*models.User, error) {
			if userID == 0 {
				return []*models.User{}, nil
			}
			return []*models.User{users[4]}, nil
		},
	}, users
}

func TestUserService_GetFriendsByUserID(t *testing.T) {
	type fields struct {
		repo IUserRepository
	}
	type args struct {
		userID int
	}
	mockRepository, users := getUserRepositoryMock()
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
				repo: mockRepository,
			},
			args:    args{userID: 1},
			want:    []*models.User{users[2], users[3]},
			wantErr: false,
		},
		{
			name: "no user",
			fields: fields{
				repo: mockRepository,
			},
			args:    args{userID: TotalUsersOfMockRepository + 1},
			want:    nil,
			wantErr: true,
		},
		{
			name: "no friends",
			fields: fields{
				repo: mockRepository,
			},
			args:    args{userID: 0},
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
	mockRepository, users := getUserRepositoryMock()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.User
		wantErr bool
	}{
		// 1-hop: (1, 2), (1, 3)
		// 2-hop: (1, 2), (1, 3), (1, 4), (1, 5)
		// block: (1, 5)
		// result: (1, 4)
		{
			name: "success",
			fields: fields{
				repo: mockRepository,
			},
			args:    args{userID: 1},
			want:    []*models.User{users[4]},
			wantErr: false,
		},
		{
			name: "no user",
			fields: fields{
				repo: mockRepository,
			},
			args:    args{userID: TotalUsersOfMockRepository + 1},
			want:    nil,
			wantErr: true,
		},
		{
			name: "no friends",
			fields: fields{
				repo: mockRepository,
			},
			args:    args{userID: 0},
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
	mockRepository, users := getUserRepositoryMock()
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
				repo: mockRepository,
			},
			args:    args{userID: 1, page: 1, limit: 5},
			want:    []*models.User{users[4]},
			wantErr: false,
			assertCalls: func(mock *mocks.IUserRepositoryMock) {
				calls := mock.FindFriendsOfFriendsExcludingSomeUsersByUserIDWithPaginationCalls()
				assert.Equal(t, 1, len(calls))
				assert.Equal(t, []int{2, 3, 5}, calls[0].ExcludedUserIDs)
			},
		},
		{
			name: "no user",
			fields: fields{
				repo: mockRepository,
			},
			args:        args{userID: TotalUsersOfMockRepository + 1, page: 1, limit: 5},
			want:        nil,
			wantErr:     true,
			assertCalls: nil,
		},
		{
			name: "no friends",
			fields: fields{
				repo: mockRepository,
			},
			args:        args{userID: 0, page: 1, limit: 5},
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
