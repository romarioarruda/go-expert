package main

import (
	"io"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"go-expert/challenges/cep-locate/structs"
)


func main() {
	fmt.Println("Informed param: ", os.Args[1:])

	for _, cep := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/"+ cep +"/json/")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error trying to make request: %v\n", err)
		}

		defer req.Body.Close()

		resp, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error trying to read response: %v\n", err)
		}

		var address viacep.ViaCep
		err = json.Unmarshal(resp, &address)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error trying to parse the response: %v\n", err)
		}

		file, err := os.Create("address.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error trying to create file address.txt: %v\n", err)
		}

		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Locale: %s, UF: %s", address.Cep, address.Localidade, address.Uf))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error trying to write into file address.txt: %v\n", err)
		}

		fmt.Println("Process finished!")
	}
}