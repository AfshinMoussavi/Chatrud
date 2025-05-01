package user

import (
	"Chat-Websocket/internal/db"
	"Chat-Websocket/pkg/loggerPkg"
	"Chat-Websocket/pkg/redisPkg"
	"Chat-Websocket/pkg/utils"
	"Chat-Websocket/pkg/validatorPkg"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type userService struct {
	userRepo  IUserRepository
	logger    loggerPkg.ILogger
	validator validatorPkg.IValidator
	timeout   time.Duration
}

func NewService(repository IUserRepository, logger loggerPkg.ILogger, validator validatorPkg.IValidator) IUserService {
	return &userService{
		userRepo:  repository,
		logger:    logger,
		validator: validator,
		timeout:   time.Duration(2) * time.Second,
	}
}

func (s *userService) CreateUserService(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	if err := s.validator.ValidateStruct(req); err != nil {
		formattedErrors := s.validator.FormatErrors(req, err)
		return nil, fmt.Errorf("validation failed: %+v", formattedErrors)
	}

	_, err := s.userRepo.GetUserByEmailRepository(ctx, req.Email)
	if err == nil {
		return &CreateUserRes{}, errors.New("user with this email already exists")
	}
	_, err = s.userRepo.GetUserByPhoneRepository(ctx, req.Phone)
	if err == nil {
		return &CreateUserRes{}, errors.New("user with this phone already exists")
	}
	_, err = s.userRepo.GetUserByNameRepository(ctx, req.Name)
	if err == nil {
		return &CreateUserRes{}, errors.New("user with this name already exists")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return &CreateUserRes{}, errors.New("hashing password failed")
	}

	u := &db.CreateUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: hashedPassword,
	}

	r, err := s.userRepo.CreateUserRepository(ctx, *u)
	if err != nil {
		return &CreateUserRes{}, fmt.Errorf("create user repository failed: %v", err)
	}

	/*
		Clear the cache
	*/
	err = redisPkg.Rdb.Del(ctx, "users:list").Err()
	if err != nil {
		s.logger.Warn("Failed to delete old cache: %v", err)
	}

	res := &CreateUserRes{
		ID:    strconv.Itoa(int(r.ID)),
		Name:  r.Name,
		Email: r.Email,
		Phone: r.Phone,
	}

	return res, nil
}

func (s *userService) ListUserService(ctx context.Context) (*[]CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	const redisKey = "users:list"

	/*
		Step 1: Try fetching from redis
	*/
	cached, err := redisPkg.Rdb.Get(ctx, redisKey).Result()
	if err == nil {
		var cachedUsers []CreateUserRes
		if err := json.Unmarshal([]byte(cached), &cachedUsers); err == nil {
			s.logger.Info("data load from redis")
			return &cachedUsers, nil
		}
	}

	/*
		Step 2: Fetch from database
	*/
	dbUsers, err := s.userRepo.ListUserRepository(ctx)
	if err != nil {
		return nil, fmt.Errorf("list user repository failed: %v", err)
	}

	var users []CreateUserRes
	for _, u := range dbUsers {
		users = append(users, CreateUserRes{
			ID:    strconv.Itoa(int(u.ID)),
			Name:  u.Name,
			Email: u.Email,
			Phone: u.Phone,
		})
	}

	/*
		Step 3: Save to Redis
	*/
	data, err := json.Marshal(users)
	if err != nil {
		s.logger.Warn("Failed to marshal users for redis: %v", err)
	} else {
		err = redisPkg.Rdb.Set(ctx, redisKey, string(data), s.timeout).Err()
		err = redisPkg.Rdb.Set(ctx, redisKey, string(data), 10*time.Minute).Err()
		if err != nil {
			s.logger.Warn("Failed to save users in redis: %v", err)
		}
	}
	s.logger.Info("data load from database")

	return &users, nil
}

func (s *userService) LoginUserService(ctx context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	email := req.Email
	user, err := s.userRepo.GetUserByEmailRepository(ctx, email)
	if err != nil {
		return &LoginUserRes{}, fmt.Errorf("get user by email failed: %v", err)
	}

	err = utils.CheckPassword(req.Password, user.Password)
	if err != nil {
		return &LoginUserRes{}, fmt.Errorf("password check failed: %v", err)
	}

	res := &LoginUserRes{
		ID:   strconv.Itoa(int(user.ID)),
		Name: user.Name,
	}

	return res, nil
}

func (s *userService) UpdateUserService(ctx context.Context, req *EditUserReq) (*EditUserRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	user, err := s.userRepo.GetUserByIdRepository(ctx, req.ID)
	if err != nil {
		return nil, fmt.Errorf("cannot fetch current user: %w", err)
	}

	name := user.Name
	if req.Name != nil {
		name = *req.Name
	}
	email := user.Email
	if req.Email != nil {
		email = *req.Email
	}
	phone := user.Phone
	if req.Phone != nil {
		phone = *req.Phone
	}

	params := &db.UpdateUserParams{
		ID:    req.ID,
		Name:  name,
		Email: email,
		Phone: phone,
	}

	updatedUser, err := s.userRepo.UpdateUserRepository(ctx, *params)
	if err != nil {
		return &EditUserRes{}, fmt.Errorf("update user repository failed: %v", err)
	}

	res := &EditUserRes{
		ID:    strconv.Itoa(int(updatedUser.ID)),
		Name:  updatedUser.Name,
		Email: updatedUser.Email,
		Phone: updatedUser.Phone,
	}

	return res, nil
}
