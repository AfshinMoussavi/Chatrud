package user_integration_test

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateUserAPI(t *testing.T) {
	router, tx := newTestRouterWithTx(t, true)
	defer tx.Rollback()

	body := `{
		"username": "sample",
		"email": "unique_sample@example.com",
		"phone": "09129876543",
		"password": "password"
	}`

	req, err := http.NewRequest(http.MethodPost, "/api/auth/register", strings.NewReader(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	t.Logf("Response: %s", resp.Body.String())
	require.Equal(t, http.StatusCreated, resp.Code)

	var res map[string]interface{}
	err = json.Unmarshal(resp.Body.Bytes(), &res)
	require.NoError(t, err)

	data, ok := res["data"].(map[string]interface{})
	require.True(t, ok)
	require.Equal(t, "sample", data["username"])
}
