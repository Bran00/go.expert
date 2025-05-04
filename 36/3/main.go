package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID	 int `gorm:"primaryKey"`
	Name string
	Products []Product
}

type Product struct {
	ID    		 		int `gorm:"primaryKey"`
	Name  		 		string
	Price 		 		float64
	CategoryID 		int
	Category 	 		Category
	SerialNumber 	SerialNumber
	gorm.Model
}

type SerialNumber struct{
	ID 						int `gorm:"primaryKey"`
	SerialNumber 	string
	ProductID 		int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	/* // create category
	category := Category{Name: "Furniture"}
	db.Create(&category)

	// create product
	product := Product{
		Name: "Cabinet", 
		Price: 258.00, 
		CategoryID: category.ID,
	}
	db.Create(&product) 

	// create serial number
	db.Create(&SerialNumber{
		SerialNumber: "1234567890",
		ProductID: product.ID,
	}) */

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println("Category:", category.Name)
		for _, product := range category.Products {
			println("Product:", product.Name, "Serial Number:", product.SerialNumber.SerialNumber)
		}
	}	
}