package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func relations() {
	db, err := gorm.Open("postgres", "postgres://postgres@database?sslmode=disable")
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	fmt.Println("connected the database")
	defer db.Close()

	// Migrate the schema
	db.DropTableIfExists(&Product{})
	db.DropTableIfExists(&Order{})
	fmt.Println("dropped tables")
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Order{})
	fmt.Println("created tables")

	// Create
	products := []Product{
		Product{
			Code: "L1212",
			Price: 1000,
		},
		Product{
			Code: "L1213",
			Price: 2000,
		},
	}
	for _, p := range products {
		db.Create(&p)
	}

	db.Find(&products)
	orders := []Order{
		Order{
			Status: "paid",
			Product: products[0],
		},
		Order{
			Status: "cancelled",
			Product: products[1],
		},
		Order{
			Status: "unpaid",
			Product: products[1],
		},
	}
	for _, o := range orders {
		db.Create(&o)
	}

	var joinedOrders []Order
	db.Preload("Product").Find(&joinedOrders)
	for _, o := range orders {
		fmt.Println(o.Status)
		fmt.Printf("\t%s\n", o.Product.Code)
	}

	var joinedProducts []Product
	db.Preload("Orders").Find(&joinedProducts)
	for _, p := range joinedProducts {
		fmt.Println(p.Code)
		for _, o := range p.Orders {
			fmt.Printf("\t%s\n", o.Status)
		}
	}
}
