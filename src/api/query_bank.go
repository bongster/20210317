package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	sdkCfg "github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/x/bank/client/cli"
)

// QueryBankTotalHandler return back total supply value from query
// referrence by cmd/fxcored/cmd/root.go:68
func (server *Server) QueryBankTotalHandler(w http.ResponseWriter, r *http.Request) {
	cmd := cli.GetCmdQueryTotalSupply()
	initClientCtx := client.Context{}
	initClientCtx, err := client.ReadPersistentCommandFlags(initClientCtx, cmd.Flags())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"err": err.Error()})
		return
	}

	initClientCtx, err = sdkCfg.ReadFromClientConfig(initClientCtx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"err": err.Error()})
		return
	}
	if err := client.SetCmdClientContextHandler(initClientCtx, cmd); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"err": err.Error()})
		return
	}
	fmt.Println(cmd.Context())
	if v := cmd.Context().Value(client.ClientContextKey); v != nil {
		clientCtxPtr := v.(*client.Context)
		fmt.Println(clientCtxPtr)
	}
	fmt.Println(cmd.Context().Value(client.ClientContextKey))
	// clientCtx, err := client.GetClientQueryContext(cmd)
	// clientCtx.Output = w
	// clientCtx.WithNodeURI("https://fx-json.functionx.io:26657")
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadGateway)
	// 	json.NewEncoder(w).Encode(map[string]string{"err": err.Error()})
	// 	return
	// }
	// queryClient := types.NewQueryClient(clientCtx)
	// res, err := queryClient.TotalSupply(nil, &types.QueryTotalSupplyRequest{})
	// clientCtx.PrintProto(res)
}
