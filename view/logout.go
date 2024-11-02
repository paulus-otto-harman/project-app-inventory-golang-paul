package view

import (
	"context"
	"database/sql"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
	handle "project/handler"
	"project/model"
	"project/util"
	com "project/util"
)

type Logout struct {
	Id          string `json:"id"`
	IsCancelled bool   `json:"-"`
	Confirmed   bool   `json:"-"`
}

func (logout *Logout) Render(ctx context.Context, db *sql.DB) int {
	com.H1("Logout")
	logout.Id = ctx.Value("sessionId").(string)
	util.PrintJson(logout)

	fmt.Printf("\nUpdate body.json to logout. Session ID is %v\n\n", ctx.Value("sessionId"))

	if answer := com.ContinueOrReturn(); answer == 0 {
		logout.IsCancelled = true
		return 0
	}

	util.PrintJson(util.ReadJson(&Logout{}, "body").(*Logout), "Request")
	handle.Logout(db)
	response := util.ReadJson(&model.Response{}, "response").(*model.Response)
	util.PrintJson(response, "Response")

	switch response.StatusCode {
	case 200:
		gola.Wait("Press any key to continue")
		logout.Confirmed = true
		return 0
	case 401:
		gola.Wait(gola.Tf(gola.Color, "Logout Failed", gola.LightYellow))
	}
	return -1
}
