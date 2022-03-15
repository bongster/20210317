package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHealthCheck(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthcheck", nil)
	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)
	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)
	t.Logf("%v", body)
	require.NotZero(t, body)
	require.Equal(t, body["ok"], true)
}
