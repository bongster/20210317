package api

import "log"

const serverAddress = "0.0.0.0:8080"

func Run() {
	server := NewServer()
	err := server.Start(serverAddress)
	if err != nil {
		log.Fatalln("can not start server:", err)
	}
}
