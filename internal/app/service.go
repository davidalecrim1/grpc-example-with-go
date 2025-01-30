package app

import "github.com/google/uuid"

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
