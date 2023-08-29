package main

import (
	"encoding/json"
	"io"
	"fmt"
	"net/http"
)

type User struct {
	Id int	`json:"id"`
	Login string `json:"login"`
}

func main() {
	req, err := http.Get("https://api.github.com/users/romarioarruda")
	if err != nil {
		panic(err)
	}

	defer req.Body.Close() // Defer will close stream connection after all process ahead runned

	resp, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	var gitUser User
	err = json.Unmarshal(resp, &gitUser) //Transform json in a acessible Struct or object
	if err != nil {
		panic(err)
	}

	fmt.Println("Resp: ", gitUser)
	fmt.Println("gitUser login: ", gitUser.Login)
}