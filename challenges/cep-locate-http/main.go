package main

import (
	"encoding/json"
	"io"
	"net/http"
	"go-expert/challenges/cep-locate/structs"
)

func main() {
	http.HandleFunc("/cep", HandleCep)

	http.ListenAndServe(":8080", nil)
}

func HandleCep(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")

	cepParam := req.URL.Query().Get("number")
	if cepParam == "" {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(`{"msg": "Param 'number' not informed"}`))
		return
	}

	result, err := FindCep(cepParam)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"msg": "There was an error to get cep"}`))
		return
	}
	
	// finalJson, err := json.Marshal(result) //Convert struct in Json value
	// if err != nil {
	// 	resp.WriteHeader(http.StatusInternalServerError)
	// 	resp.Write([]byte(`{"msg": "There was an error trying to convert json result"}`))
	// 	return
	// }

	// resp.Write(finalJson)
	json.NewEncoder(resp).Encode(result) //Convert struct in json data
}

func FindCep(cep string) (*viacep.ViaCep, error) {
	req, err := http.Get("https://viacep.com.br/ws/"+ cep +"/json/")
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var cepResult viacep.ViaCep

	err = json.Unmarshal(body, &cepResult) //Transform json in an acessible Struct
	if err != nil {
		return nil, err
	}

	return &cepResult, nil
}