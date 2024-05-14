package entity

import (
	"errors"
	"time"

	"github.com/antoniomjr/go/9-apis/pkg/entity"
)

var (
	ErrIDIsRequired = errors.New("di is required")
	ErrIvalidID = errors.New("invalid id")
	ErrNameRequired = errors.New("name is required")
	ErrPriceRequired = errors.New("price is required")
	ErrorInvalidPrice = errors.New("invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64       `json:"price"`
	CreatedAt time.Time    `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID: entity.NewID(),
		Name: name,
		Price: price,
		CreatedAt: time.Now(),
	}
	if err := product.Validate(); err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrIvalidID
	}
	if p.Name == "" {	
		return ErrNameRequired
	}
	if p.Price == 0 {
		return ErrPriceRequired
	}
	if p.Price < 0 {	
		return ErrorInvalidPrice
	}
	return nil
}
