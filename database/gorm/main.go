package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/driver/mysql"
)

type Category struct {
	ID int `gorm:"primaryKey"`
	Name string
}

type User struct {
	ID	  int `gorm:"primaryKey"`
	Name  string
	Email string
	CategoryID int
	Category Category
	gorm.Model
}

func main() {
	dsn := "root:r@@t@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{}, &Category{})
	fmt.Println("Migration finished successfuly")

	//simple create
	// category := Category{ Name: "Admin"}
	// db.Create(&category)
	// db.Create(&User{
	// 	CategoryID: 1,
	// 	Name: "Romário Dev 2",
	// 	Email: "romariodev2@gmail.com",
	// })

	//create in batch
	// db.Create([]User{
	// 	{
	// 		ID: 2,
	// 		Name: "Romário Dev 2",
	// 		Email: "romariodev2@gmail.com",
	// 	},
	// 	{
	// 		ID: 3,
	// 		Name: "Romário Dev 3",
	// 		Email: "romariodev3@gmail.com",
	// 	},
	// })

	//Find first record by id
	// var user User
	// db.First(&user, 1)
	// fmt.Println(user)

	//Find first by where
	// db.First(&user, "email = ?", "romariodev3@gmail.com")
	// fmt.Println(user)

	// var users []User
	//select all / select all using where
	// db.Find(&users, "email LIKE ?", "%romariodev%")
	// db.Where("email LIKE ?", "%romariodev2%").Find(&users)
	// db.Find(&users)
	// for _, rowUser := range users {
	// 	fmt.Println(rowUser)
	// }

	//select users and relationship categories using Belongs to technique
	// var users []User
	// db.Preload("Category").Find(&users)
	// for _, rowUser := range users {
	// 	fmt.Printf("User: %v, Category: %v\n", rowUser.Name, rowUser.Category.Name)
	// }

	//Update record
	// var user User
	// db.First(&user, 1)
	// user.Name = "Romario Dev 1"
	// db.Save(&user)

	//Delete record
	// var user2 User
	// db.First(&user2, 1)
	// db.Delete(&user2)
	// fmt.Println("User deleted: ", user2.ID)

	//Transaction / Pessimistic lock
	transaction := db.Begin()
	var category Category
	err = transaction.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&category, 2).Error
	if err != nil {
		panic(err)
	}

	category.Name = "Operator"
	transaction.Debug().Save(&category)
	transaction.Commit()

	//Default way to perform transaction
	var user User
	db.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&user, "email = ?", "romariodev@gmail.com")
}
