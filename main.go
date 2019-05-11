package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	fmt.Println("===== CRUD operations =====")
	crud()
	fmt.Println("===== Relations =====")
	relations()
}
