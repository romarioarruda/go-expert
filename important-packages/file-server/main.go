package main

import (
	"log"
	"net/http"
)

func main() {
	//In a real world application is preferrable use a web server like nginx to serve static content
	fileServer := http.FileServer(http.Dir("./important-packages/file-server/public/"))
	mux := http.NewServeMux()

	mux.Handle("/", fileServer)
	mux.HandleFunc("/test", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("Test 123"))
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}