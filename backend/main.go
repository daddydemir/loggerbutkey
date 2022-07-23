package main

import (
	"loggerbutkey/handlers"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    ":1337",
		Handler: handlers.MainRouting(),
	}
	server.ListenAndServe()
}
