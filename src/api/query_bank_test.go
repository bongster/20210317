package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	"github.com/stretchr/testify/require"
)

func TestBankModuleContext(t *testing.T) {
	cmd := cli.GetCmdQueryTotalSupply()
	clientCtx, err := client.GetClientQueryContext(cmd)
	require.NoError(t, err)
	require.NotEmpty(t, clientCtx)
}

func TestQuerybankHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/query/bank/total", nil)
	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)
	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)
	fmt.Println(body)
	require.NotZero(t, body)
	require.Equal(t, body["ok"], true)
}
