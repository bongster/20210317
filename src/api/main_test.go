package api

import (
	"net/http"
	"testing"
)

var (
	testRouter *http.ServeMux
)

func TestMain(m *testing.M) {
	testServer := NewServer()
	testRouter = testServer.router
}
