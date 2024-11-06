package main

import (
	"context"
	"project/app"
	"project/config"
	"project/database"
	view "project/view/transaction"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go app.Container(&wg, config.SessionTimeout)
	wg.Wait()
}

func test() {
	db := database.DbOpen()
	defer db.Close()
	const test = "491f841a-0e4b-40bb-b998-cae61f1568f6"
	(view.TransactionIndex{}).Render(context.WithValue(context.Background(), "sessionId", test), db)
}
