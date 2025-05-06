package user_test

import (
	"Chat-Websocket/internal/db"
	"Chat-Websocket/internal/user"
	mockLogger "Chat-Websocket/tests/unit/mocks/logger"
	mockRedis "Chat-Websocket/tests/unit/mocks/redis"
	mockUser "Chat-Websocket/tests/unit/mocks/user"
	mockValidator "Chat-Websocket/tests/unit/mocks/validator"
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestListUserService_RedisSetError(t *testing.T) {
	repo := mockUser.NewMockUserRepository(t)
	logger := mockLogger.NewMockLogger(t)
	redisClient := mockRedis.NewMockRedis(t)
	validator := mockValidator.NewMockValidator(t)

	svc := user.NewService(repo, logger, validator, redisClient)

	dbUsers := []db.User{
		{
			ID:        1,
			Name:      "Ali",
			Email:     "ali@example.com",
			Phone:     "09123456789",
			Password:  "hashed_password",
			Active:    true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	redisCmd := &redis.StringCmd{}
	redisCmd.SetErr(redis.Nil)
	redisClient.On("Get", mock.Anything, "users:list").Return(redisCmd)

	repo.On("ListUserRepository", mock.Anything).Return(dbUsers, nil)

	redisErr := errors.New("redis set error")
	redisStatusCmd := &redis.StatusCmd{}
	redisStatusCmd.SetErr(redisErr)
	redisClient.On("Set", mock.Anything, "users:list", mock.Anything, 10*time.Minute).Return(redisStatusCmd)

	logger.On("Warn", []interface{}{"Failed to save users in redis: %v", redisErr}).Return(nil)
	logger.On("Info", []interface{}{"data load from database"}).Return(nil)

	result, err := svc.ListUserService(context.Background())

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, *result, 1)
	assert.Equal(t, "Ali", (*result)[0].Name)
	assert.Equal(t, "1", (*result)[0].ID)

	redisClient.AssertExpectations(t)
	repo.AssertExpectations(t)
	logger.AssertCalled(t, "Warn", []interface{}{"Failed to save users in redis: %v", redisErr})
	logger.AssertCalled(t, "Info", []interface{}{"data load from database"})
}

func TestListUserService_SuccessFromRedis(t *testing.T) {
	repo := mockUser.NewMockUserRepository(t)
	logger := mockLogger.NewMockLogger(t)
	redisClient := mockRedis.NewMockRedis(t)
	validator := mockValidator.NewMockValidator(t)

	svc := user.NewService(repo, logger, validator, redisClient)

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

	cachedData, _ := json.Marshal(users)

	redisCmd := &redis.StringCmd{}
	redisCmd.SetVal(string(cachedData))
	redisClient.On("Get", mock.Anything, "users:list").Return(redisCmd)

	logger.On("Info", []interface{}{"data load from redis"}).Return(nil)

	result, err := svc.ListUserService(context.Background())

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, *result, 2)
	assert.Equal(t, "Ali", (*result)[0].Name)
	assert.Equal(t, "1", (*result)[0].ID)
	assert.Equal(t, "Reza", (*result)[1].Name)
	assert.Equal(t, "2", (*result)[1].ID)

	redisClient.AssertExpectations(t)
	logger.AssertCalled(t, "Info", []interface{}{"data load from redis"})
	repo.AssertNotCalled(t, "ListUserRepository", mock.Anything)
	redisClient.AssertNotCalled(t, "Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything)

}

func TestListUserService_EmptyList(t *testing.T) {
	// ایجاد موک‌ها
	repo := mockUser.NewMockUserRepository(t)
	logger := mockLogger.NewMockLogger(t)
	redisClient := mockRedis.NewMockRedis(t)
	validator := mockValidator.NewMockValidator(t)

	svc := user.NewService(repo, logger, validator, redisClient)

	dbUsers := []db.User{}

	redisCmd := &redis.StringCmd{}
	redisCmd.SetErr(redis.Nil)
	redisClient.On("Get", mock.Anything, "users:list").Return(redisCmd)

	repo.On("ListUserRepository", mock.Anything).Return(dbUsers, nil)

	redisStatusCmd := &redis.StatusCmd{}
	redisClient.On("Set", mock.Anything, "users:list", "null", 10*time.Minute).Return(redisStatusCmd)

	logger.On("Info", []interface{}{"data load from database"}).Return(nil)

	result, err := svc.ListUserService(context.Background())

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, *result, 0)

	redisClient.AssertExpectations(t)
	repo.AssertExpectations(t)
	logger.AssertCalled(t, "Info", []interface{}{"data load from database"})
	logger.AssertNotCalled(t, "Warn", mock.Anything)
}
