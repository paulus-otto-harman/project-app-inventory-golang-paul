package service

import (
	"project/model"
	"project/repository"
)

type UserService struct {
	User repository.User
}

func InitUserService(repo repository.User) *UserService {
	return &UserService{User: repo}
}

func (repo *UserService) Create(user model.User) *model.Response {
	err := repo.User.Create(&user)

	if err != nil {
		return &model.Response{StatusCode: 401, Message: "Fail to create user", Data: err.Error()}
	}
	return &model.Response{StatusCode: 200, Message: "User Created", Data: user}
}
