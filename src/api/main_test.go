package api

import (
	"testing"

	"github.com/gorilla/mux"
)

var (
	testRouter *mux.Router
)

func TestMain(m *testing.M) {
	testServer := NewServer()
	testRouter = testServer.router
}
