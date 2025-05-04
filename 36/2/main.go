package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID	 		int 		`gorm:"primaryKey"`
	Name	 	string
	Price	 	float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// create
	/* db.Create(&Product{
		Name: "Alienware",
		Price: 2500.00,
	}) */

	// create batch
	/* products := []Product{
		{Name: "IdeaPad", Price: 1500.00},
		{Name: "MacBook", Price: 3000.00},
		{Name: "Surface", Price: 2000.00},
	}
	db.Create(&products) */

	// select one
	/* var product Product
	/* db.First(&product, 1)
	fmt.Println(product)
	db.First(&product, "name = ?", "MacBook")
	fmt.Println(product) */

	// select all
	var products []Product
	db.Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
}