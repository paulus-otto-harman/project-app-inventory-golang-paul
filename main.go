package main

import (
	"context"
	"project/app"
	"project/config"
	"project/database"
	"project/view/location"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}

	wg.Add(1)
	go app.Container(&wg, config.SessionTimeout)
	wg.Wait()
}

func xmain() {
	db := database.DbOpen()
	(view.Location{}).Render(context.WithValue(context.Background(), "sessionId", "9f547358-4d44-4c21-8937-7f3985c5d8bc"), db)
}
