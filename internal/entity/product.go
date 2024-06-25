package entity

import (
	"ApiGolang/pkg/entity"
	"errors"
	"time"
)

var (
	ErrIDIsRequired    = errors.New("id is required")
	ErrinvalidID       = errors.New("invalid id")
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrinvalidPrice    = errors.New("invalid price")
)

type Product struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	CreateAt string    `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:       entity.NewID(),
		Name:     name,
		Price:    price,
		CreateAt: time.Now().String(),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrinvalidID
	}

	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price == 0.0 {
		return ErrPriceIsRequired
	}

	if p.Price < 0.0 {
		return ErrinvalidPrice
	}

	return nil

}
