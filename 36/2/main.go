package main

import (
	//"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
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
	/* var products []Product
	db.Find(&products)
	for _, product := range products {
		fmt.Println(product)
	} */

	// pagination and limit
	/* var products []Product
	db.Limit(2).Offset(2).Find(&products)
	for _, product := range products {
		fmt.Println(product)
	} */

	// where
	/* var products []Product
	db.Where("price > ?", 2000.00).Find(&products)
	for _, product := range products {
		fmt.Println(product)
	} */

	// like
	/* var products []Product
	db.Where("name Like ?", "%e%").Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}  */

	// update e deleting
	/* var p Product
	db.First(&p, 1)
	p.Name = "Alienware X17"
	db.Save(&p) 
	
		var p2 Product
		db.First(&p2, 1)
		//fmt.Println(p2.Name)

		db.Delete(&p2)
 */
}
