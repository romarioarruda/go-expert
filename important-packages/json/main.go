package main

import (
	"os"
	"encoding/json"
)

type User struct {
	Name string
	Id string
}

func main() {
	user := User{ Id: "123", Name: "Romario Dev" }

	resp, err := json.Marshal(user) //Convert object or struct in Json value
	if err != nil {
		panic(err)
	}

	println(string(resp))

	err = json.NewEncoder(os.Stdout).Encode(user) //Another way for json data return
	if err != nil {
		panic(err)
	}
}