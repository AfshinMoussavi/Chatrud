package user_test

import (
	"Chat-Websocket/internal/user"
	mockLogger "Chat-Websocket/tests/unit/mocks/logger"
	mockUser "Chat-Websocket/tests/unit/mocks/user"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userService := mockUser.NewMockUserService(t)

	logger := mockLogger.NewMockLogger(t)

	handler := user.NewHandler(userService, logger)

	users := []user.CreateUserRes{
		{
			ID:    "1",
			Name:  "Ali",
			Email: "ali@example.com",
			Phone: "09123456789",
		},
		{
			ID:    "2",
			Name:  "Reza",
			Email: "reza@example.com",
			Phone: "09123456788",
		},
	}

	userService.On("ListUserService", mock.Anything).Return(&users, nil)

	req, err := http.NewRequest(http.MethodGet, "/api/auth/users", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	router := gin.Default()
	router.GET("/api/auth/users", handler.ListUserHandler)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response user.ListUserRes
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Len(t, response.Data, 2)
	assert.Equal(t, users[0].Name, response.Data[0].Name)
	assert.Equal(t, users[1].Email, response.Data[1].Email)
}

func TestListUsersError_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userService := mockUser.NewMockUserService(t)
	logger := mockLogger.NewMockLogger(t)

	handler := user.NewHandler(userService, logger)

	userService.On("ListUserService", mock.Anything).Return(nil, errors.New("database error"))

	logger.On("Error", mock.Anything, mock.Anything).Return(nil).Maybe()

	req, err := http.NewRequest(http.MethodGet, "/api/auth/users", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	router := gin.Default()
	router.GET("/api/auth/users", handler.ListUserHandler)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Contains(t, response, "error", "Response should contain an 'error' key")
}

func TestListUsers_EmptyList(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userService := mockUser.NewMockUserService(t)
	logger := mockLogger.NewMockLogger(t)

	handler := user.NewHandler(userService, logger)

	emptyUsers := []user.CreateUserRes{}
	userService.On("ListUserService", mock.Anything).Return(&emptyUsers, nil)

	req, err := http.NewRequest(http.MethodGet, "/api/auth/users", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	router := gin.Default()
	router.GET("/api/auth/users", handler.ListUserHandler)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response user.ListUserRes
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response.Data, 0)
}
