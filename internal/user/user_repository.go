package user

import (
	"Chat-Websocket/internal/db"
	"context"
	"errors"
	"fmt"
)

type userRepository struct {
	queries *db.Queries
}

func NewRepository(queries *db.Queries) IUserRepository {
	return &userRepository{queries}
}

func (u *userRepository) CreateUserRepository(ctx context.Context, user db.CreateUserParams) (db.User, error) {
	createdUser, err := u.queries.CreateUser(ctx, db.CreateUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: user.Password,
	})

	if err != nil {
		return db.User{}, errors.New("user creation failed")
	}
	return createdUser, nil
}

func (u *userRepository) ListUserRepository(ctx context.Context) ([]db.User, error) {
	users, err := u.queries.ListUsers(ctx)
	if err != nil {
		return []db.User{}, errors.New("user list failed")
	}
	return users, nil
}

func (u *userRepository) GetUserByEmailRepository(ctx context.Context, email string) (db.User, error) {
	user, err := u.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return db.User{}, errors.New("user get by email failed")
	}
	return user, nil
}

func (u *userRepository) GetUserByPhoneRepository(ctx context.Context, phone string) (db.User, error) {
	user, err := u.queries.GetUserByPhone(ctx, phone)
	if err != nil {
		return db.User{}, errors.New("user get by phone failed")
	}
	return user, nil
}

func (u *userRepository) GetUserByNameRepository(ctx context.Context, name string) (db.User, error) {
	user, err := u.queries.GetUserByName(ctx, name)
	if err != nil {
		return db.User{}, errors.New("user get by name failed")
	}
	return user, nil
}

func (u *userRepository) UpdateUserRepository(ctx context.Context, user db.UpdateUserParams) (db.User, error) {
	updatedUser, err := u.queries.UpdateUser(ctx, user)
	if err != nil {
		return db.User{}, errors.New("user update failed")
	}
	return updatedUser, nil
}

func (u *userRepository) GetUserByIdRepository(ctx context.Context, id int32) (db.User, error) {
	user, err := u.queries.GetUserByID(ctx, id)
	if err != nil {
		return db.User{}, errors.New("user get by id failed")
	}
	return user, nil
}

func (u *userRepository) DeleteUserRepository(ctx context.Context, id int32) error {
	err := u.queries.DeleteUser(ctx, id)
	fmt.Println("err is", err)
	if err != nil {
		return err
	}
	return nil
}
