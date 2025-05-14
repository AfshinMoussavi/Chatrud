package user_integration_test

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListUsersAPI(t *testing.T) {
	router := newTestRouter(t, false)

	req, _ := http.NewRequest(http.MethodGet, "/api/auth/users", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)
}
