package app

import (
	"fmt"

	"github.com/google/uuid"
)

type Product struct {
	Id   string
	Name string
}

func NewProduct(name string) *Product {
	id := uuid.New()
	return &Product{
		Id:   id.String(),
		Name: name,
	}
}

type ProductService struct {
	products []*Product
}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) Add(p *Product) {
	s.products = append(s.products, p)
}

func (s *ProductService) Delete(id string) {
	for i, p := range s.products {
		if p.Id == id {
			s.products = append(s.products[:i], s.products[i+1:]...)
			break
		}
	}
}

func (s *ProductService) Get(id string) (*Product, error) {
	for _, p := range s.products {
		if p.Id == id {
			return p, nil
		}
	}

	return nil, fmt.Errorf("product not found")
}

func (s *ProductService) Update(u *Product) {
	for i, p := range s.products {
		if p.Id == u.Id {
			s.products[i] = u
			break
		}
	}
}
