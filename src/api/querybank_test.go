package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/require"
)

const (
	defaultFxGrpcUrl = "localhost:9090"
)

func TestQueryBankTotal(t *testing.T) {
	clientConn, err := grpcNewClient(defaultFxGrpcUrl)
	require.NoError(t, err)
	bankQueryClient := banktypes.NewQueryClient(clientConn)
	res, err := bankQueryClient.TotalSupply(context.Background(), &banktypes.QueryTotalSupplyRequest{})
	require.NoError(t, err)
	require.NotEmpty(t, res)
}

func TestQueryBankTotalHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/query/bank/total", nil)
	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)
	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)
	t.Logf("%v", body)
	require.Equal(t, body["ok"], true)
}
