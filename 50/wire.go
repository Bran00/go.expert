//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/Bran00/go.expert/49/product"
	"github.com/google/wire"
)

var setRepositoryDepedency = wire.NewSet(
	product.NewProductRepository,
	wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)),
)

func NewUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		setRepositoryDepedency,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}