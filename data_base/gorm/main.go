package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

// func NewProduct(name string, price float64) *Product {
// 	return &Product{
// 		ID:    uuid.New().String(),
// 		Name:  name,
// 		Price: price,
// 	}
// }

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
