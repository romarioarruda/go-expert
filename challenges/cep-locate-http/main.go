package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/cep", FindCep)

	http.ListenAndServe(":8080", nil)
}

func FindCep(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")

	cepParam := req.URL.Query().Get("number")
	if cepParam == "" {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(`{"msg": "Param 'number' not informed"}`))
		return
	}

	resp.Write([]byte(`{"msg": "Hello World"}`))
}