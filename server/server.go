package server

import (
	"fmt"
	"net/http"
)

func Server() {
	router := http.NewServeMux()
	NewHelloHandler(router)

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	fmt.Println("http://localhost:3000")
	server.ListenAndServe()
}
