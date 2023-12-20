package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World !!"))
	})

	fmt.Printf("Server listen at http://localhost:9000")
	http.ListenAndServe(":9000", nil)
}
