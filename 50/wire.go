//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/Bran00/go.expert/49/product"
	"github.com/google/wire"
)

func NewUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		product.NewProductRepository,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}