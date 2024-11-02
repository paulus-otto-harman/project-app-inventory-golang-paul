package service

import (
	"project/model"
	"project/repository"
)

type LocationService struct {
	Location repository.Location
}

func InitLocationService(repo repository.Location) *LocationService {
	return &LocationService{Location: repo}
}

func (repo *LocationService) Create(location model.Location) *model.Response {
	err := repo.Location.Create(&location)

	if err != nil {
		return &model.Response{StatusCode: 401, Message: "Fail to create Location", Data: err.Error()}
	}
	return &model.Response{StatusCode: 200, Message: "Location Created", Data: location}
}
