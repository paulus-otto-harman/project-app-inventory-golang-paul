package handler

import (
	"database/sql"
	"project/model"
	"project/repository"
	"project/service"
	"project/util"
)

func Login(db *sql.DB) {
	// input
	user := util.ReadJson(&model.User{}, "body").(*model.User)

	// proses
	result := service.InitAuthService(*repository.CreateAuth(db)).Login(*user)

	// output
	util.BuildJson(result, "response")
}
