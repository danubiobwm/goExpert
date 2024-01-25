package main

import (
	"database/sql"

	"github.com/danubiobwm/goExpert/di/product"
	"github.com/google/wire"
)

var setRepositoryDependencyWire = wire.NewSet(product.NewProductRepository,
	wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)))

func NewUseCaseWire(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		product.NewProductUseCase,
		setRepositoryDependencyWire,
	)
	return &product.ProductUseCase{}
}
