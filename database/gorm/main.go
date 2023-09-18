package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

type User struct {
	ID	  int `gorm:"primaryKey"`
	Name  string
	Email string
	gorm.Model
}

func main() {
	dsn := "root:r@@t@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})

	//simple create
	// db.Create(&User{
	// 	ID: 1,
	// 	Name: "Romário Dev",
	// 	Email: "romariodev@gmail.com",
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

	//Update record
	// var user User
	// db.First(&user, 1)
	// user.Name = "Romario Dev 1"
	// db.Save(&user)

	//Delete record
	var user2 User
	db.First(&user2, 1)
	db.Delete(&user2)
	fmt.Println("User deleted: ", user2.ID)
}
