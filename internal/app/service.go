package app

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

var ErrProductNotFound = fmt.Errorf("product not found")

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
	mu       sync.RWMutex
}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) Add(p *Product) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.products = append(s.products, p)
}

func (s *ProductService) Delete(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, p := range s.products {
		if p.Id == id {
			s.products = append(s.products[:i], s.products[i+1:]...)
			break
		}
	}
}

func (s *ProductService) Get(id string) (*Product, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, p := range s.products {
		if p.Id == id {
			return p, nil
		}
	}

	return nil, ErrProductNotFound
}

func (s *ProductService) Update(u *Product) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, p := range s.products {
		if p.Id == u.Id {
			s.products[i] = u
			break
		}
	}
}
