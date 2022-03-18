package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	Name               = "test"
	defaultNodeGrpcUrl = "https://fx-json.functionx.io:26657"
)

func newHttpClient(rpcUrl string) (*rpchttp.HTTP, error) {
	return rpchttp.New(rpcUrl, "/websocket")
}

func grpcNewClient(grpcUrl string) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	// http use  grpc.WithInsecure()
	// opts = append(opts, grpc.WithInsecure())
	// https use grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "*.function.io"))
	opts = append(opts, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, strings.Split(grpcUrl, ":")[0])))
	return grpc.Dial(grpcUrl, opts...)
}

// QueryBankTotalHandler return back total supply value from query
// referrence by cmd/fxcored/cmd/root.go:68
func (server *Server) QueryBankTotalHandler(w http.ResponseWriter, r *http.Request) {
	// Create dummy cobra.Command
	cmd := &cobra.Command{}
	// Calling Execute function because if not calling this function then cmd.Context are nil
	cmd.Execute()
	rpcURI := defaultNodeGrpcUrl
	cmd.Flags().Set(flags.FlagNode, rpcURI)
	clientCtx, err := client.GetClientQueryContext(cmd)
	clientCtx.WithOutput(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"err": err.Error()})
		return
	}

	queryClient := types.NewQueryClient(clientCtx)
	res, err := queryClient.TotalSupply(cmd.Context(), &types.QueryTotalSupplyRequest{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"err": err.Error()})
		return
	}
	clientCtx.PrintProto(res)
}
