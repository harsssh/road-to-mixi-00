package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"problem1/mocks"
	"problem1/models"
	"problem1/services"
	"testing"
)

func TestUserHandler_GetFriendList(t *testing.T) {
	type fields struct {
		service services.IUserService
	}
	type response struct {
		status    int
		checkBody bool
		body      string
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func() (*http.Request, *httptest.ResponseRecorder)
		want    response
	}{
		{
			name: "success",
			fields: fields{
				service: &mocks.IUserServiceMock{
					GetFriendListFunc: func(userID int) ([]*models.User, error) {
						return []*models.User{
							{UserID: 2, Name: "User2"},
							{UserID: 3, Name: "User3"},
						}, nil
					},
				},
			},
			prepare: func() (*http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodGet, "/get_friend_list?id=1", nil)
				rec := httptest.NewRecorder()
				return req, rec
			},
			want: response{
				status:    http.StatusOK,
				checkBody: true,
				body:      `[{"user_id":2,"name":"User2"},{"user_id":3,"name":"User3"}]`,
			},
		},
		{
			name: "no id",
			fields: fields{
				service: &mocks.IUserServiceMock{},
			},
			prepare: func() (*http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodGet, "/get_friend_list", nil)
				rec := httptest.NewRecorder()
				return req, rec
			},
			want: response{
				status:    http.StatusBadRequest,
				checkBody: false,
			},
		},
		{
			name: "invalid id",
			fields: fields{
				service: &mocks.IUserServiceMock{},
			},
			prepare: func() (*http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodGet, "/get_friend_list?id=invalid", nil)
				rec := httptest.NewRecorder()
				return req, rec
			},
			want: response{
				status:    http.StatusBadRequest,
				checkBody: false,
			},
		},
		{
			name: "negative id",
			fields: fields{
				service: &mocks.IUserServiceMock{},
			},
			prepare: func() (*http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodGet, "/get_friend_list?id=-1", nil)
				rec := httptest.NewRecorder()
				return req, rec
			},
			want: response{
				status:    http.StatusBadRequest,
				checkBody: false,
			},
		},
		{
			name: "no user",
			fields: fields{
				service: &mocks.IUserServiceMock{
					GetFriendListFunc: func(userID int) ([]*models.User, error) {
						return nil, services.ErrUserNotFound
					},
				},
			},
			prepare: func() (*http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodGet, "/get_friend_list?id=1", nil)
				rec := httptest.NewRecorder()
				return req, rec
			},
			want: response{
				status:    http.StatusNotFound,
				checkBody: false,
			},
		},
		{
			name: "no friends",
			fields: fields{
				service: &mocks.IUserServiceMock{
					GetFriendListFunc: func(userID int) ([]*models.User, error) {
						return []*models.User{}, nil
					},
				},
			},
			prepare: func() (*http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodGet, "/get_friend_list?id=1", nil)
				rec := httptest.NewRecorder()
				return req, rec
			},
			want: response{
				status:    http.StatusOK,
				checkBody: true,
				body:      `[]`,
			},
		},
		{
			name: "service error",
			fields: fields{
				service: &mocks.IUserServiceMock{
					GetFriendListFunc: func(userID int) ([]*models.User, error) {
						return nil, fmt.Errorf("error")
					},
				},
			},
			prepare: func() (*http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodGet, "/get_friend_list?id=1", nil)
				rec := httptest.NewRecorder()
				return req, rec
			},
			want: response{
				status:    http.StatusInternalServerError,
				checkBody: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &UserHandler{
				service: tt.fields.service,
			}
			e := echo.New()
			req, rec := tt.prepare()
			c := e.NewContext(req, rec)
			require.NoError(t, h.GetFriendList(c))
			assert.Equal(t, tt.want.status, rec.Code)
			if tt.want.checkBody {
				assert.JSONEq(t, tt.want.body, rec.Body.String())
			}
		})
	}
}

func TestUserHandler_GetFriendOfFriendList(t *testing.T) {
	type fields struct {
		service services.IUserService
	}
	type response struct {
		status    int
		checkBody bool
		body      string
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func() (*http.Request, *httptest.ResponseRecorder)
		want    response
	}{
		{
			name: "success",
			fields: fields{
				service: &mocks.IUserServiceMock{
					GetFriendOfFriendListFunc: func(userID int) ([]*models.User, error) {
						return []*models.User{
							{UserID: 2, Name: "User2"},
							{UserID: 3, Name: "User3"},
						}, nil
					},
				},
			},
			prepare: func() (*http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodGet, "/get_friend_of_friend_list?id=1", nil)
				rec := httptest.NewRecorder()
				return req, rec
			},
			want: response{
				status:    http.StatusOK,
				checkBody: true,
				body:      `[{"user_id":2,"name":"User2"},{"user_id":3,"name":"User3"}]`,
			},
		},
		{
			name: "no id",
			fields: fields{
				service: &mocks.IUserServiceMock{},
			},
			prepare: func() (*http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodGet, "/get_friend_of_friend_list", nil)
				rec := httptest.NewRecorder()
				return req, rec
			},
			want: response{
				status:    http.StatusBadRequest,
				checkBody: false,
			},
		},
		{
			name: "invalid id",
			fields: fields{
				service: &mocks.IUserServiceMock{},
			},
			prepare: func() (*http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodGet, "/get_friend_of_friend_list?id=invalid", nil)
				rec := httptest.NewRecorder()
				return req, rec
			},
			want: response{
				status:    http.StatusBadRequest,
				checkBody: false,
			},
		},
		{
			name: "negative id",
			fields: fields{
				service: &mocks.IUserServiceMock{},
			},
			prepare: func() (*http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodGet, "/get_friend_of_friend_list?id=-1", nil)
				rec := httptest.NewRecorder()
				return req, rec
			},
			want: response{
				status:    http.StatusBadRequest,
				checkBody: false,
			},
		},
		{
			name: "no user",
			fields: fields{
				service: &mocks.IUserServiceMock{
					GetFriendOfFriendListFunc: func(userID int) ([]*models.User, error) {
						return nil, services.ErrUserNotFound
					},
				},
			},
			prepare: func() (*http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodGet, "/get_friend_of_friend_list?id=1", nil)
				rec := httptest.NewRecorder()
				return req, rec
			},
			want: response{
				status:    http.StatusNotFound,
				checkBody: false,
			},
		},
		{
			name: "no friends",
			fields: fields{
				service: &mocks.IUserServiceMock{
					GetFriendOfFriendListFunc: func(userID int) ([]*models.User, error) {
						return []*models.User{}, nil
					},
				},
			},
			prepare: func() (*http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodGet, "/get_friend_of_friend_list?id=1", nil)
				rec := httptest.NewRecorder()
				return req, rec
			},
			want: response{
				status:    http.StatusOK,
				checkBody: true,
				body:      `[]`,
			},
		},
		{
			name: "service error",
			fields: fields{
				service: &mocks.IUserServiceMock{
					GetFriendOfFriendListFunc: func(userID int) ([]*models.User, error) {
						return nil, fmt.Errorf("error")
					},
				},
			},
			prepare: func() (*http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodGet, "/get_friend_of_friend_list?id=1", nil)
				rec := httptest.NewRecorder()
				return req, rec
			},
			want: response{
				status:    http.StatusInternalServerError,
				checkBody: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &UserHandler{
				service: tt.fields.service,
			}
			e := echo.New()
			req, rec := tt.prepare()
			c := e.NewContext(req, rec)
			require.NoError(t, h.GetFriendOfFriendList(c))
			assert.Equal(t, tt.want.status, rec.Code)
			if tt.want.checkBody {
				assert.JSONEq(t, tt.want.body, rec.Body.String())
			}
		})
	}
}

func TestUserHandler_GetFriendOfFriendListPaging(t *testing.T) {
	type fields struct {
		service services.IUserService
	}
	type response struct {
		status    int
		checkBody bool
		body      string
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func() (*http.Request, *httptest.ResponseRecorder)
		want    response
	}{
		{
			name: "success",
			fields: fields{
				service: &mocks.IUserServiceMock{
					GetFriendOfFriendListPagingFunc: func(userID int, page int, limit int) ([]*models.User, error) {
						return []*models.User{
							{UserID: 2, Name: "User2"},
							{UserID: 3, Name: "User3"},
						}, nil
					},
				},
			},
			prepare: func() (*http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(
					http.MethodGet,
					"/get_friend_of_friend_list_paging?id=1&page=1&limit=3",
					nil,
				)
				rec := httptest.NewRecorder()
				return req, rec
			},
			want: response{
				status:    http.StatusOK,
				checkBody: true,
				body:      `[{"user_id":2,"name":"User2"},{"user_id":3,"name":"User3"}]`,
			},
		},
		{
			name: "no id",
			fields: fields{
				service: &mocks.IUserServiceMock{},
			},
			prepare: func() (*http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(
					http.MethodGet,
					"/get_friend_of_friend_list_paging?page=1&limit=3",
					nil,
				)
				rec := httptest.NewRecorder()
				return req, rec
			},
			want: response{
				status:    http.StatusBadRequest,
				checkBody: false,
			},
		},
		{
			name: "invalid id",
			fields: fields{
				service: &mocks.IUserServiceMock{},
			},
			prepare: func() (*http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(
					http.MethodGet,
					"/get_friend_of_friend_list_paging?id=invalid&page=1&limit=3",
					nil,
				)
				rec := httptest.NewRecorder()
				return req, rec
			},
			want: response{
				status:    http.StatusBadRequest,
				checkBody: false,
			},
		},
		{
			name: "negative id",
			fields: fields{
				service: &mocks.IUserServiceMock{},
			},
			prepare: func() (*http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(
					http.MethodGet,
					"/get_friend_of_friend_list_paging?id=-1&page=1&limit=3",
					nil,
				)
				rec := httptest.NewRecorder()
				return req, rec
			},
			want: response{
				status:    http.StatusBadRequest,
				checkBody: false,
			},
		},
		{
			name: "no user",
			fields: fields{
				service: &mocks.IUserServiceMock{
					GetFriendOfFriendListPagingFunc: func(userID int, page int, limit int) ([]*models.User, error) {
						return nil, services.ErrUserNotFound
					},
				},
			},
			prepare: func() (*http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(
					http.MethodGet,
					"/get_friend_of_friend_list_paging?id=1&page=1&limit=3",
					nil,
				)
				rec := httptest.NewRecorder()
				return req, rec
			},
			want: response{
				status:    http.StatusNotFound,
				checkBody: false,
			},
		},
		{
			name: "no friends",
			fields: fields{
				service: &mocks.IUserServiceMock{
					GetFriendOfFriendListPagingFunc: func(userID int, page int, limit int) ([]*models.User, error) {
						return []*models.User{}, nil
					},
				},
			},
			prepare: func() (*http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(
					http.MethodGet,
					"/get_friend_of_friend_list_paging?id=1&page=1&limit=3",
					nil,
				)
				rec := httptest.NewRecorder()
				return req, rec
			},
			want: response{
				status:    http.StatusOK,
				checkBody: true,
				body:      `[]`,
			},
		},
		{
			name: "service error",
			fields: fields{
				service: &mocks.IUserServiceMock{
					GetFriendOfFriendListPagingFunc: func(userID int, page int, limit int) ([]*models.User, error) {
						return nil, fmt.Errorf("error")
					},
				},
			},
			prepare: func() (*http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(
					http.MethodGet,
					"/get_friend_of_friend_list_paging?id=1&page=1&limit=3",
					nil,
				)
				rec := httptest.NewRecorder()
				return req, rec
			},
			want: response{
				status:    http.StatusInternalServerError,
				checkBody: false,
			},
		},
		{
			name: "default page and limit",
			fields: fields{
				service: &mocks.IUserServiceMock{
					GetFriendOfFriendListPagingFunc: func(userID int, page int, limit int) ([]*models.User, error) {
						return []*models.User{}, nil
					},
				},
			},
			prepare: func() (*http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(
					http.MethodGet,
					"/get_friend_of_friend_list_paging?id=1",
					nil,
				)
				rec := httptest.NewRecorder()
				return req, rec
			},
			want: response{
				status:    http.StatusOK,
				checkBody: false,
			},
		},
		{
			name: "page 0",
			fields: fields{
				service: &mocks.IUserServiceMock{},
			},
			prepare: func() (*http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(
					http.MethodGet,
					"/get_friend_of_friend_list_paging?id=1&page=0&limit=3",
					nil,
				)
				rec := httptest.NewRecorder()
				return req, rec
			},
			want: response{
				status:    http.StatusBadRequest,
				checkBody: false,
			},
		},
		{
			name: "limit 0",
			fields: fields{
				service: &mocks.IUserServiceMock{},
			},
			prepare: func() (*http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(
					http.MethodGet,
					"/get_friend_of_friend_list_paging?id=1&page=1&limit=0",
					nil,
				)
				rec := httptest.NewRecorder()
				return req, rec
			},
			want: response{
				status:    http.StatusBadRequest,
				checkBody: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &UserHandler{
				service: tt.fields.service,
			}
			e := echo.New()
			req, rec := tt.prepare()
			c := e.NewContext(req, rec)
			require.NoError(t, h.GetFriendOfFriendListPaging(c))
			assert.Equal(t, tt.want.status, rec.Code)
			if tt.want.checkBody {
				assert.JSONEq(t, tt.want.body, rec.Body.String())
			}
		})
	}
}
