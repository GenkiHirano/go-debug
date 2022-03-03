package model

import (
	"errors"
	"time"
)

type Product struct {
	ID        int
	Name      string
	CreatedAt time.Time
}

func NewProduct(id int, name string) (*Product, error) {
	if id == 0 {
		return nil, errors.New("idを入力してください")
	}

	if name == "" {
		return nil, errors.New("nameを入力してください")
	}

	now := time.Now()

	product := &Product{
		ID:        id,
		Name:      name,
		CreatedAt: now,
	}

	return product, nil
}
