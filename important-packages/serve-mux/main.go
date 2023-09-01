package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/test", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("Hello mux"))
	})

	http.ListenAndServe(":8080", mux) //replacing default multiplexer
}