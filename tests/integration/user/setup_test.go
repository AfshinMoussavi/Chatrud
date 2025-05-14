package user_integration_test

import (
	"Chat-Websocket/internal/db"
	"Chat-Websocket/internal/user"
	mocksLogger "Chat-Websocket/tests/unit/mocks/logger"
	mocksRedis "Chat-Websocket/tests/unit/mocks/redis"
	mocksValidator "Chat-Websocket/tests/unit/mocks/validator"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func newTestDB(t *testing.T) *db.Queries {
	connStr := "postgresql://postgres:1234@localhost:5432/test_chat_db?sslmode=disable"
	database, err := sql.Open("postgres", connStr)
	require.NoError(t, err)
	return db.New(database)
}

func newTestRouter(t *testing.T, forCreateUser bool) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	logger := mocksLogger.NewMockLogger(t)
	validator := mocksValidator.NewMockValidator(t)
	redisMock := mocksRedis.NewMockRedis(t)

	if forCreateUser {
		logger.On("Warn", mock.Anything, mock.Anything).Return().Maybe()
		redisMock.On("Del", mock.Anything, mock.Anything).Return(&redis.IntCmd{}, nil).Maybe()
		validator.On("ValidateStruct", mock.Anything).Return(nil).Maybe()
	} else {
		logger.On("Info", mock.Anything).Return().Maybe()
		redisMock.On("Get", mock.Anything, mock.Anything).Return(&redis.StringCmd{}, nil).Maybe()
		redisMock.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&redis.StatusCmd{}, nil).Maybe()
	}

	repo := user.NewRepository(newTestDB(t))
	service := user.NewService(repo, logger, validator, redisMock)
	handler := user.NewHandler(service, logger)

	r.GET("/api/auth/users", handler.ListUserHandler)
	r.POST("/api/auth/register", handler.CreateUserHandler)

	return r
}

func newTestDBWithTx(t *testing.T) (*sql.Tx, *db.Queries) {
	connStr := "postgresql://postgres:1234@localhost:5432/test_chat_db?sslmode=disable"
	dbConn, err := sql.Open("postgres", connStr)
	require.NoError(t, err)

	tx, err := dbConn.Begin()
	require.NoError(t, err)

	queries := db.New(tx)
	return tx, queries
}

func newTestRouterWithTx(t *testing.T, forCreateUser bool) (*gin.Engine, *sql.Tx) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	logger := mocksLogger.NewMockLogger(t)
	validator := mocksValidator.NewMockValidator(t)
	redisMock := mocksRedis.NewMockRedis(t)

	if forCreateUser {
		logger.On("Warn", mock.Anything, mock.Anything).Return().Maybe()
		redisMock.On("Del", mock.Anything, mock.Anything).Return(&redis.IntCmd{}, nil).Maybe()
		validator.On("ValidateStruct", mock.Anything).Return(nil).Maybe()
	} else {
		logger.On("Info", mock.Anything).Return().Maybe()
		redisMock.On("Get", mock.Anything, mock.Anything).Return(&redis.StringCmd{}, nil).Maybe()
		redisMock.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&redis.StatusCmd{}, nil).Maybe()
	}

	tx, queries := newTestDBWithTx(t)
	repo := user.NewRepository(queries)
	service := user.NewService(repo, logger, validator, redisMock)
	handler := user.NewHandler(service, logger)

	r.GET("/api/auth/users", handler.ListUserHandler)
	r.POST("/api/auth/register", handler.CreateUserHandler)

	return r, tx
}
