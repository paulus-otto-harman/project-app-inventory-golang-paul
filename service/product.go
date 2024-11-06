package service

import (
	"project/model"
	"project/repository"
)

type ProductService struct {
	Product repository.Product
}

func InitProductService(repo repository.Product) *ProductService {
	return &ProductService{Product: repo}
}

func (repo *ProductService) Create(product model.Product) *model.Response {
	err := repo.Product.Create(&product)

	if err != nil {
		return &model.Response{StatusCode: 401, Message: "Fail to create Product", Data: err.Error()}
	}
	return &model.Response{StatusCode: 200, Message: "Product successfully created", Data: product}
}

func (repo *ProductService) Update(product model.Product) *model.Response {
	err := repo.Product.Update(&product)

	if err != nil {
		return &model.Response{StatusCode: 401, Message: "Fail to update Product", Data: err.Error()}
	}
	return &model.Response{StatusCode: 200, Message: "Product's Balance successfully updated", Data: struct {
		Id      int `json:"id"`
		Balance int `json:"balance"`
	}{Id: product.Id, Balance: product.Balance}}
}

func (repo *ProductService) Search(criteria model.Search) *model.DataPage {
	page, limit, totalItems, totalPages, products, err := repo.Product.Search(criteria)

	if err != nil {
		return &model.DataPage{StatusCode: 404, Message: "Search failed", Data: err.Error()}
	}
	return &model.DataPage{StatusCode: 200, Message: "Products Found", Page: page, Limit: limit, TotalItems: totalItems, TotalPages: totalPages, Data: products}
}
