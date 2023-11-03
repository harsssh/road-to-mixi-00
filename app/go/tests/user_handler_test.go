package tests

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"problem1/handlers"
	"problem1/models"
	"testing"
)

func TestUserHandler_GetFriendList(t *testing.T) {
	tests := []struct {
		name           string
		data           []*models.User
		query          string
		expectedStatus int
		expectedBody   string
		checkBody      bool
	}{
		{
			name:           "no param",
			data:           nil,
			query:          "",
			expectedStatus: http.StatusBadRequest,
			checkBody:      false,
		},
		{
			name:           "invalid id",
			data:           nil,
			query:          "id=a",
			expectedStatus: http.StatusBadRequest,
			checkBody:      false,
		},
		{
			name:           "found",
			data:           []*models.User{{UserID: 1, Name: "A"}, {UserID: 2, Name: "B"}},
			query:          "id=1",
			expectedStatus: http.StatusOK,
			expectedBody:   `[{"user_id":1,"name":"A"},{"user_id":2,"name":"B"}]`,
			checkBody:      true,
		},
	}

	e := echo.New()
	for _, tt := range tests {
		req := httptest.NewRequest(echo.GET, "/get_friend_list", nil)
		req.URL.RawQuery = tt.query
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		t.Run(tt.name, func(t *testing.T) {
			h := handlers.NewUserHandler(NewMockUserService(tt.data))
			err := h.GetFriendList(c)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.expectedStatus, rec.Code)
			if tt.checkBody {
				assert.JSONEq(t, tt.expectedBody, rec.Body.String())
			}
		})
	}
}
