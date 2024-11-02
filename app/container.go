package app

import (
	"context"
	"project/database"
	"project/view"
	"sync"
	"time"
)

func sessionTimeout(timeout int) time.Duration {
	if timeout == 0 {
		return time.Until(time.Now().AddDate(100, 0, 0))
	}
	return time.Duration(timeout) * time.Second
}

func Container(wg *sync.WaitGroup, timeout int) {
	defer wg.Done()

	db := database.DbOpen()
	defer db.Close()

	sessionLifetime := sessionTimeout(timeout)
	for {
		loginScreen := view.Login{}
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), sessionLifetime)
			defer cancel()

			view.Render(&loginScreen, ctx, db)
		}()

		if loginScreen.IsCancelled {
			return
		}

		//fmt.Println(loginScreen.SessionId)

	}
}
