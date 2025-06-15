//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Bran00/go.expert/49/product"
	"github.com/google/wire"
)

func NewUseCase() *product.ProductUseCase {
	wire.Build(
		product.NewProductRepository,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}