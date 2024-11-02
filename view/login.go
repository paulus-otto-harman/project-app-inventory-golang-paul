package view

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
	handle "project/handler"
	"project/model"
	"project/util"
	com "project/util"
)

type Login struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	IsCancelled bool   `json:"-"`
}

func (login *Login) Render(ctx context.Context, db *sql.DB) int {
	com.H1("Login")

	util.PrintJson(login)
	fmt.Printf("\nUpdate body.json to login\n\n")

	if answer := com.ContinueOrExit(); answer == 0 {
		login.IsCancelled = true
		return 0
	}

	util.PrintJson(util.ReadJson(&Login{}, "body").(*Login), "Request")
	handle.Login(db)
	response := util.ReadJson(&model.Response{}, "response").(*model.Response)
	util.PrintJson(response, "Response")

	switch response.StatusCode {
	case 200:
		gola.Wait("Press any key to continue")
		home := Home{}
		Render(&home, context.WithValue(ctx, "sessionId", getSessionId(response)), db)
		if !home.isLogout {
			gola.Wait("Session expired. Press Enter to login")
		}
	case 401:
		gola.Wait(gola.Tf(gola.Color, "User Not Found", gola.LightYellow))
	}

	return -1
}

func getSessionId(response *model.Response) string {
	session := model.Session{}
	sessionData, _ := json.Marshal(response.Data)
	json.Unmarshal(sessionData, &session)

	return session.Id
}
