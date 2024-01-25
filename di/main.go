package main

import (
	"database/sql"
	"fmt"
	"github.com/danubiobwm/goExpert/di/product"

	_"github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	// Create a new product repository

	repository := product.NewProductRepository(db)

	// Create a new UseCase Product
	usecase := product.NewProductUseCase(repository)

	product, err := usecase.GetProduct(1)

	if err != nil {
		panic(err)
	}

	// Print the product
	fmt.Println(product.Name)
}
