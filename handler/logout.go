package handler

import (
	"database/sql"
	"project/model"
	"project/repository"
	"project/service"
	"project/util"
)

func Logout(db *sql.DB) {
	// input
	session := util.ReadJson(&model.Session{}, "body").(*model.Session)

	// proses
	result := service.InitAuthService(*repository.CreateAuth(db)).Logout(*session)

	// output
	util.BuildJson(result, "response")
}
