package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories"`
}
type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// category := Category{Name: "Eletronicos"}
	// db.Create(&category)
	// category2 := Category{Name: "Cozinha"}
	// db.Create(&category2)
	// db.Create(&Product{
	// 	Name:       "Panela",
	// 	Price:      2103.12,
	// 	Categories: []Category{category, category2},
	// })

	var categories []Category
	errCatg := db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if errCatg != nil {
		panic(err)
	}
	for _, category := range categories {
		for _, product := range category.Products {
			println(product.Name, category.Name)
		}
	}

}