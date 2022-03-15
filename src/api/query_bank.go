package api

import (
	"encoding/json"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
)

// QueryBankTotalHandler return back total supply value from query
func (server *Server) QueryBankTotalHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	clientCtx, err := client.GetClientQueryContext(nil)
	clientCtx.Output = w
	clientCtx.WithNodeURI("https://fx-json.functionx.io:26657")
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(map[string]string{"err": err.Error()})
		return
	}
	queryClient := types.NewQueryClient(clientCtx)
	res, err := queryClient.TotalSupply(nil, &types.QueryTotalSupplyRequest{})
	clientCtx.PrintProto(res)
}
