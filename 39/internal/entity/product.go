package entity

import (
	"errors"
	"time"

	"github.com/Bran00/go.expert/39/pkg/entity"
)

var (
	ErrIDIsRequiered    = errors.New("ID is required")
	ErrInvalidID        = errors.New("invalid ID")
	ErrNameIsRequiered  = errors.New("name is required")
	ErrPriceIsRequiered = errors.New("price is required")
	ErrInvalidPrice     = errors.New("invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequiered
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidID
	}
	if p.Name == "" {
		return ErrNameIsRequiered
	}
	if p.Price == 0 {
		return ErrPriceIsRequiered
	}
	if p.Price < 0 {
		return ErrInvalidPrice
	}
	return nil
}
