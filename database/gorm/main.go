package main

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

type User struct {
	ID	  int `gorm:"primaryKey"`
	Name  string
	Email string
}

func main() {
	dsn := "root:r@@t@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})

	//create simple way
	db.Create(User{
		ID: 1,
		Name: "Romário Dev",
		Email: "romariodev@gmail.com",
	})

	//create in batch
	db.Create([]User{
		{
			ID: 2,
			Name: "Romário Dev 2",
			Email: "romariodev2@gmail.com",
		},
		{
			ID: 3,
			Name: "Romário Dev 3",
			Email: "romariodev3@gmail.com",
		},
	})
}