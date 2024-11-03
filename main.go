package main

import (
	"context"
	"project/app"
	"project/config"
	"project/database"
	"project/view/product"
	"sync"
)

func xmain() {

	wg := sync.WaitGroup{}

	wg.Add(1)
	go app.Container(&wg, config.SessionTimeout)
	wg.Wait()
}

func main() {
	db := database.DbOpen()
	(view.Product{}).Render(context.WithValue(context.Background(), "sessionId", "9f547358-4d44-4c21-8937-7f3985c5d8bc"), db)
}
