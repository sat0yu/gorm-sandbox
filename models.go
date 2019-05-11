package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
	Orders []Order
}

type Order struct {
	gorm.Model
	Status string
	ProductID int
	Product Product
}

