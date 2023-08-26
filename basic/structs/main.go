package main

import "fmt"

type Contact struct {
	phone string
	email string
}

type Client struct {
	name string
	age int
	isActive bool
	Contacts Contact
}

func (customer *Client) Deactivate() {
	customer.isActive = false
}

func main() {
	client := Client{
		name: "Romário Dev",
		age: 27,
		isActive: true,
		Contacts: Contact{
			email: "romario@dev",
			phone: "666",
		},
	}

	fmt.Println("Customer name: ", client.name)
	fmt.Println("Customer age: ", client.age)
	fmt.Println("Customer status: ", client.isActive)
	fmt.Println("Customer contacts: ", client.Contacts)

	client.Deactivate()
	fmt.Println("Customer new status: ", client.isActive)
}