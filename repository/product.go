package repository

import "github.com/GenkiHirano/go-debug/model"

func Create(id int, name string) (*model.Product, error) {
	product, err := model.NewProduct(id, name)
	if err != nil {
		return nil, err
	}

	newProduct := &model.Product{
		ID:        product.ID,
		CreatedAt: product.CreatedAt,
	}

	return newProduct, nil
}
