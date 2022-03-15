package api

import (
	"log"
	"net/http"
)

type Server struct {
	router *http.ServeMux
}

func NewServer() *Server {
	server := &Server{}
	router := http.NewServeMux()
	router.HandleFunc("/healthcheck", server.HealthCheckHandler)
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	log.Printf("Server is running %s\n", address)
	return http.ListenAndServe(address, server.router)
}
