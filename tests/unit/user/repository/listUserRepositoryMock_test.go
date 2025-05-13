package repository

import (
	"Chat-Websocket/internal/db"
	"Chat-Websocket/internal/user"
	mocks "Chat-Websocket/tests/unit/mocks/db"
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestListUsersRepository_Success(t *testing.T) {
	mockQuerier := mocks.NewMockQuerier(t)

	expectedUsers := []db.User{
		{ID: 1, Name: "Ali", Email: "ali@example.com", Phone: "09150118812"},
		{ID: 2, Name: "Sara", Email: "sara@example.com", Phone: "09150118813"},
	}

	mockQuerier.On("ListUsers", mock.Anything).Return(expectedUsers, nil)

	repo := user.NewRepository(mockQuerier)

	ctx := context.Background()
	users, err := repo.ListUserRepository(ctx)

	require.NoError(t, err)
	require.Equal(t, expectedUsers, users)

	mockQuerier.AssertExpectations(t)

}

func TestGetUserByEmailRepository_Success(t *testing.T) {
	mockQuerier := mocks.NewMockQuerier(t)
	expectedUser := db.User{ID: 1, Name: "Ali", Email: "ali@example.com", Phone: "09150118812"}

	email := "ali@example.com"
	mockQuerier.On("GetUserByEmail", mock.Anything, email).Return(expectedUser, nil)

	repo := user.NewRepository(mockQuerier)
	ctx := context.Background()
	usr, err := repo.GetUserByEmailRepository(ctx, email)

	require.NoError(t, err)
	require.Equal(t, expectedUser, usr)
	mockQuerier.AssertExpectations(t)

}

func TestCreateUserRepository_Success(t *testing.T) {
	mockQuerier := mocks.NewMockQuerier(t)

	input := db.CreateUserParams{
		Name:     "Afshin",
		Email:    "afshin@example.com",
		Phone:    "09150118812",
		Password: "password",
	}

	expectedUser := db.User{
		ID:       1,
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
		Password: input.Password,
	}

	mockQuerier.On("CreateUser", mock.Anything, input).Return(expectedUser, nil)

	repo := user.NewRepository(mockQuerier)

	ctx := context.Background()
	createdUser, err := repo.CreateUserRepository(ctx, input)

	require.NoError(t, err)
	require.Equal(t, expectedUser, createdUser)
	mockQuerier.AssertExpectations(t)
}

func TestCreateUserRepository_Error(t *testing.T) {
	mockQuerier := mocks.NewMockQuerier(t)

	input := db.CreateUserParams{
		Name:     "Afshin",
		Email:    "afshin@example.com",
		Phone:    "09150000000",
		Password: "securepassword",
	}

	mockQuerier.On("CreateUser", mock.Anything, input).Return(db.User{}, errors.New("db error"))

	repo := user.NewRepository(mockQuerier)

	ctx := context.Background()
	createdUser, err := repo.CreateUserRepository(ctx, input)

	require.Error(t, err)
	require.Equal(t, db.User{}, createdUser)
	require.EqualError(t, err, "user creation failed")

	mockQuerier.AssertExpectations(t)
}

func TestUpdateUserRepository_Success(t *testing.T) {
	mockQuerier := mocks.NewMockQuerier(t)

	input := db.UpdateUserParams{
		ID:    1,
		Name:  "Ali Updated",
		Email: "ali.updated@example.com",
		Phone: "09150118899",
	}

	expectedUser := db.User{
		ID:       1,
		Name:     "Ali Updated",
		Email:    "ali.updated@example.com",
		Phone:    "09150118899",
		Password: "newpassword",
	}

	mockQuerier.On("UpdateUser", mock.Anything, input).Return(expectedUser, nil)

	repo := user.NewRepository(mockQuerier)
	ctx := context.Background()

	actualUser, err := repo.UpdateUserRepository(ctx, input)

	require.NoError(t, err)
	require.Equal(t, expectedUser, actualUser)

	mockQuerier.AssertExpectations(t)
}

func TestUpdateUserRepository_Error(t *testing.T) {
	mockQuerier := mocks.NewMockQuerier(t)

	input := db.UpdateUserParams{
		ID:    1,
		Name:  "Ali Error",
		Email: "ali.error@example.com",
		Phone: "09150118800",
	}

	mockQuerier.On("UpdateUser", mock.Anything, input).Return(db.User{}, errors.New("db error"))

	repo := user.NewRepository(mockQuerier)
	ctx := context.Background()

	actualUser, err := repo.UpdateUserRepository(ctx, input)

	require.Error(t, err)
	require.EqualError(t, err, "user update failed")
	require.Equal(t, db.User{}, actualUser)

	mockQuerier.AssertExpectations(t)
}
