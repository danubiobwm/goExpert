package main

import (
	"database/sql"

	"github.com/danubiobwm/goExpert/di/product"
	"github.com/google/wire"
)

var setRepositoryDependency = wire.NewSet(product.NewProductRepository,
	wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)))

func NewUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		product.NewProductUseCase,
		setRepositoryDependency,
	)
	return &product.ProductUseCase{}
}
