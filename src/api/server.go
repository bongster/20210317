package api

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

// NewServer return server object
func NewServer() *Server {
	router := mux.NewRouter()
	server := &Server{}
	router.HandleFunc("/healthcheck", server.HealthCheckHandler).Methods(http.MethodGet)
	router.HandleFunc("/query/bank/total", server.QueryBankTotalHandler).Methods(http.MethodGet)
	server.router = router
	return server
}

// Start running HTTP server
func (server *Server) Start(address string) error {
	log.Printf("Server is running %s\n", address)
	srv := http.Server{
		Handler: server.router,
		Addr:    address,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return srv.ListenAndServe()
}
