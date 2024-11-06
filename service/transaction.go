package service

import (
	"project/model"
	"project/repository"
)

type TransactionService struct {
	Transaction repository.Transaction
}

func InitTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{Transaction: repo}
}

func (repo *TransactionService) Create(transaction model.Transaction) *model.Response {
	err := repo.Transaction.Create(&transaction)

	if err != nil {
		return &model.Response{StatusCode: 401, Message: "Fail to create Transaction", Data: err.Error()}
	}
	return &model.Response{StatusCode: 200, Message: "Transaction Created", Data: transaction}
}

func (repo *TransactionService) All(criteria model.Search) *model.DataPage {
	page, limit, totalItems, totalPages, transactions, err := repo.Transaction.All(criteria)

	if err != nil {
		return &model.DataPage{StatusCode: 404, Message: "No transaction found", Data: err.Error()}
	}
	return &model.DataPage{StatusCode: 200, Message: "Transactions Found", Page: page, Limit: limit, TotalItems: totalItems, TotalPages: totalPages, Data: transactions}
}
