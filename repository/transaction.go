package repository

import (
	"database/sql"
	"math"
	"project/model"
)

type Transaction struct {
	Db *sql.DB
}

func CreateTransaction(db *sql.DB) *Transaction {
	return &Transaction{Db: db}
}

func (repo *Transaction) Create(transaction *model.Transaction) error {
	tx, _ := repo.Db.Begin()
	query := `UPDATE products SET balance=balance+$1 FROM sessions WHERE products.id=$2 AND sessions.id=$3 AND expired_at IS NULL`
	qty := transaction.Quantity
	if transaction.Type == "out" {
		qty = transaction.Quantity * -1
	}
	_, err := repo.Db.Exec(query, qty, transaction.ProductId, transaction.SessionId)
	if err != nil {
		query = `INSERT INTO transactions (type,product_id,quantity,note)
				SELECT $1,$2,$3,$4 FROM sessions WHERE id=$5 AND expired_at IS NULL RETURNING id`
		err = repo.Db.QueryRow(query, transaction.Type, transaction.ProductId, transaction.Quantity, transaction.Note, transaction.SessionId).Scan(&transaction.Id)
		if err != nil {
			tx.Rollback()
		}
		tx.Commit()
	}

	if err != nil {
		return err
	}
	return nil
}

func (repo *Transaction) All(criteria model.Search) (int, int, int, float64, []model.Transaction, error) {
	var count int
	queryCount := `SELECT COUNT(*) FROM transactions`
	err := repo.Db.QueryRow(queryCount).Scan(&count)

	query := `SELECT transactions.id, transactions.type, transactions.product_id , transactions.quantity, transactions.note, transactions.created_at  
				FROM transactions
				JOIN products ON transactions.product_id = products.id
				LIMIT $1
				OFFSET $2`

	offset := (criteria.Page - 1) * criteria.Limit
	rows, err := repo.Db.Query(query, criteria.Limit, offset)
	var transactions []model.Transaction
	for rows.Next() {
		var transaction model.Transaction
		if err := rows.Scan(&transaction.Id, &transaction.Type, &transaction.ProductId, &transaction.Quantity, &transaction.Note, &transaction.DateCreated); err != nil {
			return 0, 0, 0, 0, []model.Transaction{}, err
		}
		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return 0, 0, 0, 0, []model.Transaction{}, err
	}
	return criteria.Page, criteria.Limit, count, math.Ceil(float64(count) / float64(criteria.Limit)), transactions, nil
}
