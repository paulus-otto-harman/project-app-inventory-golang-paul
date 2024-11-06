package transactions

import (
	"database/sql"
	"project/model"
	"project/repository"
	"project/service"
	"project/util"
)

func TransactionIndex(db *sql.DB) {
	// input
	searchCriteria := util.ReadJson(&model.Search{}, "body").(*model.Search)

	// proses
	result := service.InitTransactionService(*repository.CreateTransaction(db)).All(*searchCriteria)

	// output
	util.BuildJson(result, "response")
}
