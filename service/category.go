package service

import (
	"project/model"
	"project/repository"
)

type CategoryService struct {
	Category repository.Category
}

func InitCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{Category: repo}
}

func (repo *CategoryService) Create(category model.Category) *model.Response {
	err := repo.Category.Create(&category)

	if err != nil {
		return &model.Response{StatusCode: 401, Message: "Fail to create Category", Data: err.Error()}
	}
	return &model.Response{StatusCode: 200, Message: "Category Created", Data: category}
}
