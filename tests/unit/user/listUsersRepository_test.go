package user_test

import (
	"Chat-Websocket/internal/db"
	mockDb "Chat-Websocket/tests/unit/mocks/db"
	"context"

	"Chat-Websocket/internal/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestListUserRepository_Success(t *testing.T) {
	// ایجاد موک Queries
	queries := &mockDb.MockQuerier{}

	// ایجاد Repository
	repo := user.NewRepository(queries)

	// داده‌های نمونه
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
		{
			ID:        2,
			Name:      "Reza",
			Email:     "reza@example.com",
			Phone:     "09123456788",
			Password:  "hashed_password",
			Active:    true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	// تنظیم رفتار موک
	queries.On("ListUsers", mock.Anything).Return(dbUsers, nil)

	// اجرای متد Repository
	result, err := repo.ListUserRepository(context.Background())

	// بررسی نتایج
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 2)
	assert.Equal(t, "Ali", result[0].Name)
	assert.Equal(t, int32(1), result[0].ID)
	assert.Equal(t, "Reza", result[1].Name)
	assert.Equal(t, int32(2), result[1].ID)

	// بررسی فراخوانی‌های موک
	queries.AssertExpectations(t)
}
