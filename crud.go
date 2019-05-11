package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func crud() {
	db, err := gorm.Open("postgres", "postgres://postgres@database?sslmode=disable")
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	fmt.Println("connected the database")
	defer db.Close()

	// Migrate the schema
	db.DropTableIfExists(&Product{})
	fmt.Println("dropped a table")
	db.AutoMigrate(&Product{})
	fmt.Println("created a table")

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})
	db.Create(&Product{Code: "L1213", Price: 2000})

	fmt.Println("Inserted two records")
	var products []Product
	db.Find(&products)
	for _, p := range products {
		fmt.Println(p)
	}

	// Read
	var product Product

	db.First(&product, "Code = ?", "L1213") // find product with code l1213
	fmt.Println("Fetched the record which has code 'L1213'")
	fmt.Println(product)

	// Update - update product's price to 3000
	db.Model(&product).Update("Price", 3000)
	fmt.Println("Updated the record which has code 'L1213'")

	db.Find(&products)
	for _, p := range products {
		fmt.Println(p)
	}

	// Delete - delete product
	db.Delete(&product)
	fmt.Println("Deleted the record which has code 'L1213'")

	db.Find(&products)
	for _, p := range products {
		fmt.Println(p)
	}
}
