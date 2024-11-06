package transactions

import (
	"database/sql"
	"project/model"
	"project/repository"
	"project/service"
	"project/util"
)

func TransactionCreate(db *sql.DB) {
	// input
	transaction := util.ReadJson(&model.Transaction{}, "body").(*model.Transaction)

	// proses
	result := service.InitTransactionService(*repository.CreateTransaction(db)).Create(*transaction)

	// output
	util.BuildJson(result, "response")
}
