package user_handler_test

import (
	"Chat-Websocket/internal/user"
	mockLogger "Chat-Websocket/tests/unit/mocks/logger"
	mockUser "Chat-Websocket/tests/unit/mocks/user"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"

	"testing"
)

func TestCreateUserHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userService := mockUser.NewMockUserService(t)
	logger := mockLogger.NewMockLogger(t)

	handler := user.NewHandler(userService, logger)

	input := user.CreateUserReq{
		Name:     "Ali",
		Email:    "ali@gmail.com",
		Phone:    "09123456789",
		Password: "password",
	}

	expectedRes := user.CreateUserRes{
		ID:    "1",
		Name:  input.Name,
		Email: input.Email,
		Phone: input.Phone,
	}

	userService.On("CreateUserService", mock.Anything, &input).Return(&expectedRes, nil)

	body, err := json.Marshal(input)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/api/auth/register", bytes.NewBuffer(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router := gin.Default()
	router.POST("/api/auth/register", handler.CreateUserHandler)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var response map[string]user.CreateUserRes
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedRes.Name, response["data"].Name)
	assert.Equal(t, expectedRes.Email, response["data"].Email)

}

func TestCreateUserHandler_BadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userService := mockUser.NewMockUserService(t)
	logger := mockLogger.NewMockLogger(t)

	handler := user.NewHandler(userService, logger)

	invalidJSON := `{"name": "Ali", "email": "ali@example.com", "unknownField": "oops"}`

	req, err := http.NewRequest(http.MethodPost, "/api/auth/register", strings.NewReader(invalidJSON))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router := gin.Default()
	router.POST("/api/auth/register", handler.CreateUserHandler)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response["error"], "Invalid input data")
}

func TestCreateUserHandler_InternalServerError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userService := mockUser.NewMockUserService(t)
	logger := mockLogger.NewMockLogger(t)

	handler := user.NewHandler(userService, logger)

	input := user.CreateUserReq{
		Name:     "Ali",
		Email:    "ali@example.com",
		Phone:    "09123456789",
		Password: "securepassword",
	}

	userService.
		On("CreateUserService", mock.Anything, &input).
		Return(nil, errors.New("something went wrong"))

	logger.On("Error", mock.Anything, mock.Anything).Return(nil).Maybe()

	body, err := json.Marshal(input)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/api/auth/register", bytes.NewBuffer(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router := gin.Default()
	router.POST("/api/auth/register", handler.CreateUserHandler)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response["error"], "Failed to create user")
}
