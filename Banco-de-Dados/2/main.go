package main

import (
	"fmt"

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

	//create
	// db.Create(&Product{
	// 	Name:  "Carro",
	// 	Price: 2000,
	// })

	// create batch
	// product := []Product{
	//   {Name: "Carro", Price: 3000},
	//   {Name: "Notebook", Price: 2000},
	//   {Name: "TV", Price: 4000},
	//   {Name: "Sofa", Price: 5000},
	// }
	// db.Create(&product)

	//Select One
	// var product Product
	// db.First(&product,2)
	// fmt.Println("Product:", product)
	// db.First(&product, "name=?", "TV")
	// fmt.Println("Product:", product)

	//Select All
	// var products []Product
	// db.Find(&products)

	// for _,product := range products {
	//   fmt.Println("Product:", product)
	// }

	//select

	// var products []Product
	// db.Limit(2).Find(&products)
	// for _,product := range products {
	//   fmt.Println("Product:", product)
	// }

	//OFFset
	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// for _,product := range products {
	//   fmt.Println("Product:", product)
	// }

	// var products []Product
	// db.Where("price>?", 3000).Find(&products)
	// for _,product := range products {
	//    fmt.Println("Product:", product)
	// }

	// var products []Product
	// db.Where("name LIKE ?", "%book%").Find(&products)
	// for _,product := range products {
	//    fmt.Println("Product:", product)
	// }

	// var p Product
	// db.First(&p, 1)
	// p.Name = "TV2"
	// db.Save(&p)

	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p2.Name)

	db.Delete(&p2)

}
