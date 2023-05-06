package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}
type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}
type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// create
	// products := []Product{
	// 	{Name: "Notebook", Price: 1120.20},
	// 	{Name: "Geladeira", Price: 3450.10},
	// 	{Name: "Microndas", Price: 450.30},
	// }
	// db.Create(&products)

	//select one
	// var product Product
	// var product2 Product
	// db.First(&product, 1) // 1 = ID
	// fmt.Println(product)
	// db.First(&product2, "name = ?", "Geladeira")
	// fmt.Println(product2)

	//select all
	// var products []Product
	// db.Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	//select all with limit // u can use "offset" to do pagination
	// var products2 []Product
	// db.Limit(2).Find(&products2)
	// for _, product := range products2 {
	// 	fmt.Println(product)
	// }

	//select with where
	// var products3 []Product
	// db.Where("price > ?", 1000).Find(&products3)
	// for _, product := range products3 {
	// 	fmt.Println(product)
	// }

	//Edit and Delete
	// var product3 Product
	// db.First(&product3, 3)
	// fmt.Println(product3)
	// product3.Name = "Ratoeira"
	// db.Save(&product3)
	// db.First(&product3, 3)
	// fmt.Println(product3)
	// db.Delete(&product3)

	//Reletionship "Belong to"
	// category := Category{Name: "Eletronicos"}
	// db.Create(&category)
	// db.Create(&Product{
	// 	Name:       "Celular",
	// 	Price:      2103.12,
	// 	CategoryID: category.ID,
	// })
	// var products4 []Product
	// db.Preload("Category").Find(&products4)
	// for _, product := range products4 {
	// 	fmt.Println(product.Name, product.Category.Name)
	// }

	//has one
	// category2 := Category{Name: "Eletronicos"}
	// db.Create(&category2)
	// db.Create(&Product{
	// 	Name:       "Celular",
	// 	Price:      2103.12,
	// 	CategoryID: category2.ID,
	// })
	// db.Create(&SerialNumber{
	// 	Number:    "123456",
	// 	ProductID: 1, // because on this has on section moment i only have 1 product
	// })
	// var products5 []Product
	// db.Preload("Category").Preload("SerialNumber").Find(&products5)
	// for _, product := range products5 {
	// 	fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	// }

	//has many
	// category4 := Category{Name: "Eletronicos"}
	// db.Create(&category4)
	// products5 := []Product{
	// 	{Name: "Notebook", Price: 1120.20, CategoryID: 1},
	// 	{Name: "Geladeira", Price: 3450.10, CategoryID: 1},
	// 	{Name: "Microndas", Price: 450.30, CategoryID: 1},
	// }
	// db.Create(&products5)
	// serialNumbers := []SerialNumber{
	// 	{Number: "123456", ProductID: 1},
	// 	{Number: "4434653", ProductID: 2},
	// 	{Number: "97454", ProductID: 3},
	// }
	// db.Create(&serialNumbers)
	// var categories []Category
	// errCatg := db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	// if errCatg != nil {
	// 	panic(err)
	// }
	// for _, category := range categories {
	// 	for _, product := range category.Products {
	// 		println(product.Name, category.Name, product.SerialNumber.Number)
	// 	}
	// }

}
