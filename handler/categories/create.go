package categories

import (
	"database/sql"
	"project/model"
	"project/repository"
	"project/service"
	"project/util"
)

func AddCategory(db *sql.DB) {
	// input
	category := util.ReadJson(&model.Category{}, "body").(*model.Category)

	// proses
	result := service.InitCategoryService(*repository.CreateCategory(db)).Create(*category)

	// output
	util.BuildJson(result, "response")
}
