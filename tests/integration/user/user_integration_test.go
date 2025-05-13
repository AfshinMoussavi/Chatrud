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
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTestDB(t *testing.T) *db.Queries {
	connStr := "postgresql://postgres:1234@localhost:5432/test_chat_db?sslmode=disable"
	database, err := sql.Open("postgres", connStr)
	require.NoError(t, err)
	return db.New(database)
}

func setupRouter(t *testing.T) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockLogger := mocksLogger.NewMockLogger(t)
	mockValidator := mocksValidator.NewMockValidator(t)
	mockRedis := mocksRedis.NewMockRedis(t)

	mockLogger.On("Info", mock.AnythingOfType("[]interface {}")).Return()

	mockRedis.On("Get", mock.Anything, "users:list").Return(&redis.StringCmd{}, nil)

	mockRedis.On("Set", mock.Anything, "users:list", mock.Anything, mock.Anything).Return(&redis.StatusCmd{}, nil)

	queries := setupTestDB(t)
	repo := user.NewRepository(queries)
	service := user.NewService(repo, mockLogger, mockValidator, mockRedis)
	handler := user.NewHandler(service, mockLogger)

	r.GET("/api/auth/users", handler.ListUserHandler)
	return r
}

func TestGetUsersAPI(t *testing.T) {
	router := setupRouter(t)

	req, _ := http.NewRequest(http.MethodGet, "/api/auth/users", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)

}
