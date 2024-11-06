package products

import (
	"database/sql"
	"project/model"
	"project/repository"
	"project/service"
	"project/util"
)

func ProductUpdate(db *sql.DB) {
	// input
	product := util.ReadJson(&model.Product{}, "body").(*model.Product)

	// proses
	result := service.InitProductService(*repository.CreateProduct(db)).Update(*product)

	// output
	util.BuildJson(result, "response")
}
