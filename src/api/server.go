package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cosmos/cosmos-sdk/x/bank"
)

type Server struct {
	router *http.ServeMux
}

// NewServer return server object
func NewServer() *Server {
	server := &Server{}
	router := http.NewServeMux()
	router.HandleFunc("/healthcheck", server.HealthCheckHandler)
	bankModule := bank.AppModuleBasic{}
	fmt.Println(bankModule)
	router.HandleFunc("/query/bank/total", server.QueryBankTotalHandler)
	server.router = router
	return server
}

// Start running HTTP server
func (server *Server) Start(address string) error {
	log.Printf("Server is running %s\n", address)
	return http.ListenAndServe(address, server.router)
}
