package categories

import (
	"database/sql"
	"project/model"
	"project/repository"
	"project/service"
	"project/util"
)

func AddLocation(db *sql.DB) {
	// input
	location := util.ReadJson(&model.Location{}, "body").(*model.Location)

	// proses
	result := service.InitLocationService(*repository.CreateLocation(db)).Create(*location)

	// output
	util.BuildJson(result, "response")
}
