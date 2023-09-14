package main

import (
	"fmt"
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ID string
	Name string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID: uuid.New().String(),
		Name: name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:r@@t@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	product := NewProduct("Notebook", 1999.9)

	err = DbInsertProduct(db, product)
	if err != nil {
		panic(err)
	}

	product.ID = "40b682fe-2862-400d-80c7-dcc13f47a435"
	product.Price = 600.0
	product.Name = "IPhone"

	err = DbUpdateProduct(db, product)
	if err != nil {
		panic(err)
	}

	result, err := DbSelectProduct(db, product.ID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("id: %v, product: %v, price: %.2f\n", result.ID, result.Name, result.Price)

	products, err := DbSelectProducts(db)
	if err != nil {
		panic(err)
	}

	for _, value := range products {
		fmt.Printf("id: %v, name: %v, price: %.2f\n", value.ID, value.Name, value.Price)
	}

	err = DbDeleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
}

func DbSelectProducts(db *sql.DB) ([]Product, error) {
	fmt.Println("Selecting all products...")

	rows, err := db.Query("select * from products")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []Product

	for rows.Next() {
		var prod Product

		err = rows.Scan(&prod.ID, &prod.Name, &prod.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, prod)
	}

	return products, nil
}

func DbSelectProduct(db *sql.DB, id string) (*Product, error) {
	fmt.Println("Selecting one product by id...")
	stmt, err := db.Prepare("select * from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var prod Product

	err = stmt.QueryRow(id).Scan(&prod.ID, &prod.Name, &prod.Price)
	if err != nil {
		return nil, err
	}

	return &prod, nil
}

func DbInsertProduct(db *sql.DB, product *Product) error {
	fmt.Println("Inserting one product...")
	stmt, err := db.Prepare("insert into products(id, name, price) values (?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}

	return nil
}

func DbUpdateProduct(db *sql.DB, product *Product) error {
	fmt.Println("Updating one product...")
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func DbDeleteProduct(db *sql.DB, id string) error {
	fmt.Println("Deleting one product by id...")
	stmt, err := db.Prepare("delete from products where id != ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}