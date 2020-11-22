package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// dsn := "root:kim24153020@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:#kim24153020@tcp(localhost:3306)/go?charset=utf8mb4&parseTime=True&loc=Local", // data source name                                                                            // auto configure based on currently MySQL version
	}), &gorm.Config{})

	// db, err := gorm.Open("mysql", "root:#kim24153020@tcp(127.0.0.1:3306)/go?charset=utf8&parseTime=True&loc=Local")

	fmt.Print(db)
	fmt.Print("------------------------")
	fmt.Print(err)
	// if err != nil {
	// 	fmt.Println("statuse: ", err)
	// }
	// defer Config.DB.Close()
	// Config.DB.AutoMigrate(&Models.Book{})

	// r := Routers.SetupRouter()
	// running
	// r.Run()
}
