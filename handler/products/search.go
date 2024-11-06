package products

import (
	"database/sql"
	"project/model"
	"project/repository"
	"project/service"
	"project/util"
)

func ProductSearch(db *sql.DB) {
	// input
	searchCriteria := util.ReadJson(&model.Search{}, "body").(*model.Search)

	// proses
	result := service.InitProductService(*repository.CreateProduct(db)).Search(*searchCriteria)

	// output
	util.BuildJson(result, "response")
}
