package main

import (
	"net/http"

	"github.com/go-clout/hadex"
)

func main() {
	server := &http.Server{
		Addr:    "localhost:2830",
		Handler: hadex.NewCore(),
	}
	server.ListenAndServe()
}
