package product

import "database/sql"

type ProductRepositoryInterface interface {
	GetProcuct(id int) (Product, error)
}

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (r *ProductRepository) GetProcuct(id int) (Product, error) {
	return Product{
		ID:   id,
		Name: "Product 1",
	}, nil
}
