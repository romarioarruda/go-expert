package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	Id int	`json:"id"`
	Login string `json:"login"`
}

func main() {
	client := http.Client{Timeout: time.Second * 5}

	fmt.Println("Perfoming GET...")
	getGithubUser("romarioarruda", client)

	fmt.Println("Perfoming POST...")
	postGithubUser(client)
}

func getGithubUser(userName string, client http.Client) {
	req, err := client.Get("https://api.github.com/users/" + userName)
	if err != nil {
		panic(err)
	}

	defer req.Body.Close() // Defer will close stream connection after all process ahead runned

	resp, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	var gitUser User
	err = json.Unmarshal(resp, &gitUser) //Transform json in an acessible Struct or object
	if err != nil {
		panic(err)
	}

	fmt.Println("Resp: ", gitUser)
	fmt.Println("gitUser login: ", gitUser.Login)
}

func postGithubUser(client http.Client) {
	url := "https://api.github.com/users/romarioarruda"
	fmt.Println("URL: ", url)
	body := []byte(`{ userName: romarioarruda }`)
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.CopyBuffer(os.Stdout, resp.Body, nil)
}