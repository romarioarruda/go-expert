package main

import "fmt"

type Customer struct {
	name string
	email string
	status bool
}

type Person interface {
	Deactivate()
}

func (customer *Customer) Deactivate() {
	customer.status = false

	fmt.Println("Deactivating to:", customer.status)
}

func Deactivation(person Person) {
	person.Deactivate()
}

func main() {
	customer := Customer {
		name: "ROmário Dev",
		email: "romariodev@gmail.com",
		status: true,
	}

	fmt.Println("Customer: ", customer)

	Deactivation(&customer) //pass the pointer reference

	fmt.Println("Customer: ", customer)
}