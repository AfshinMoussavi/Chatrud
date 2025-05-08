package user

import (
	"Chat-Websocket/internal/db"
	"context"
	"github.com/gin-gonic/gin"
)

type IUserRepository interface {
	CreateUserRepository(ctx context.Context, user db.CreateUserParams) (db.User, error)
	ListUserRepository(ctx context.Context) ([]db.User, error)
	GetUserByEmailRepository(ctx context.Context, email string) (db.User, error)
	GetUserByPhoneRepository(ctx context.Context, phone string) (db.User, error)
	GetUserByNameRepository(ctx context.Context, name string) (db.User, error)
	GetUserByIdRepository(ctx context.Context, id int32) (db.User, error)
	UpdateUserRepository(ctx context.Context, user db.UpdateUserParams) (db.User, error)
	DeleteUserRepository(ctx context.Context, id int32) error
}

type IUserService interface {
	CreateUserService(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error)
	ListUserService(ctx context.Context) (*[]CreateUserRes, error)
	LoginUserService(ctx context.Context, req *LoginUserReq) (*LoginUserRes, error)
	UpdateUserService(ctx context.Context, req *EditUserReq) (*EditUserRes, error)
	DeleteUserService(ctx context.Context, id int32) error
}

type IUserHandler interface {
	CreateUserHandler(c *gin.Context)
	ListUserHandler(c *gin.Context)
	LoginUserHandler(c *gin.Context)
	EditUserHandler(c *gin.Context)
	DeleteUserHandler(c *gin.Context)
}

type CreateUserReq struct {
	Name     string `json:"username" validate:"required,min=2,max=100,alpha"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required,numeric,len=11,mobile"`
	Password string `json:"password" validate:"required"`
}

type CreateUserRes struct {
	ID    string `json:"id"`
	Name  string `json:"username"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserRes struct {
	ID    string `json:"id"`
	Name  string `json:"username"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type EditUserReq struct {
	ID    int32   `json:"-"`
	Name  *string `json:"username" validate:"min=2,max=100,alpha"`
	Email *string `json:"email" validate:"email"`
	Phone *string `json:"phone" validate:"numeric,len=11,mobile"`
}

type EditUserRes struct {
	ID    string `json:"id"`
	Name  string `json:"username"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type ListUserRes struct {
	Data []CreateUserRes `json:"data"`
}
