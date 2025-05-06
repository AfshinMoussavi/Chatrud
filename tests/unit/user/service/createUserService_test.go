package user_service_test

import (
	"Chat-Websocket/internal/db"
	"Chat-Websocket/internal/user"

	mockLogger "Chat-Websocket/tests/unit/mocks/logger"
	mockRedis "Chat-Websocket/tests/unit/mocks/redis"
	mockUser "Chat-Websocket/tests/unit/mocks/user"
	mockValidator "Chat-Websocket/tests/unit/mocks/validator"
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateUserService_Success(t *testing.T) {
	repo := mockUser.NewMockUserRepository(t)
	logger := mockLogger.NewMockLogger(t)
	redisClient := mockRedis.NewMockRedis(t)
	validator := mockValidator.NewMockValidator(t)

	svc := user.NewService(repo, logger, validator, redisClient)

	req := &user.CreateUserReq{
		Name:     "Ali",
		Email:    "ali@example.com",
		Phone:    "09123456789",
		Password: "password",
	}

	validator.On("ValidateStruct", req).Return(nil)

	repo.On("GetUserByEmailRepository", mock.Anything, req.Email).
		Return(db.User{}, errors.New("not found"))
	repo.On("GetUserByPhoneRepository", mock.Anything, req.Phone).
		Return(db.User{}, errors.New("not found"))
	repo.On("GetUserByNameRepository", mock.Anything, req.Name).
		Return(db.User{}, errors.New("not found"))

	createdUser := db.User{
		ID:       1,
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: "hashed_password",
	}
	repo.On("CreateUserRepository", mock.Anything, mock.MatchedBy(func(p db.CreateUserParams) bool {
		return p.Email == req.Email && p.Name == req.Name && p.Phone == req.Phone
	})).Return(createdUser, nil)

	redisClient.On("Del", mock.Anything, mock.MatchedBy(func(keys interface{}) bool {
		keySlice, ok := keys.([]string)
		return ok && len(keySlice) == 1 && keySlice[0] == "users:list"
	})).Return(redis.NewIntCmd(context.Background()))

	res, err := svc.CreateUserService(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "1", res.ID)
	assert.Equal(t, req.Name, res.Name)
	assert.Equal(t, req.Email, res.Email)
	assert.Equal(t, req.Phone, res.Phone)

	repo.AssertExpectations(t)
	validator.AssertExpectations(t)
	redisClient.AssertExpectations(t)
}
