package main

import (
	"time"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/context", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	log.Println("Request started")

	select {
		case <- time.After(3 * time.Second):
			log.Println("Request processed successfully") // log into server terminal
			resp.Write([]byte("Request processed successfully")) //log into web page
		case <- ctx.Done():
			log.Println("Request was cancelled by the client side") // log into server terminal
	}

	log.Println("Request finished")
}