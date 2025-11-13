package store

import (
	"errors"
	"sync"

	"github.com/you/product-api/models"
)

var ErrNotFound = errors.New("not found")

type InMemoryStore struct {
	mu       sync.RWMutex
	products map[int64]*models.Product
	nextID   int64
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		products: make(map[int64]*models.Product),
		nextID:   1,
	}
}

func (s *InMemoryStore) Create(p *models.Product) (*models.Product, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	p.ID = s.nextID
	s.nextID++
	// copy to avoid outside modifications
	cp := *p
	s.products[p.ID] = &cp
	return &cp, nil
}

func (s *InMemoryStore) Get(id int64) (*models.Product, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if p, ok := s.products[id]; ok {
		cp := *p
		return &cp, nil
	}
	return nil, ErrNotFound
}

func (s *InMemoryStore) List() ([]models.Product, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]models.Product, 0, len(s.products))
	for _, p := range s.products {
		out = append(out, *p)
	}
	return out, nil
}

func (s *InMemoryStore) Update(id int64, upd *models.Product) (*models.Product, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if p, ok := s.products[id]; ok {
		// update fields (except ID)
		p.Name = upd.Name
		p.Description = upd.Description
		p.Price = upd.Price
		p.Stock = upd.Stock
		cp := *p
		return &cp, nil
	}
	return nil, ErrNotFound
}

func (s *InMemoryStore) Delete(id int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.products[id]; ok {
		delete(s.products, id)
		return nil
	}
	return ErrNotFound
}
