package view

import (
	"context"
	"database/sql"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
	handle "project/handler/locations"
	"project/model"
	"project/util"
	com "project/util"
)

type Location struct {
	Id        int    `json:"-"`
	Name      string `json:"name"`
	SessionId string `json:"session_id"`
}

func (location Location) Render(ctx context.Context, db *sql.DB) int {
	com.H1("Lokasi Penyimpanan")
	location.SessionId = ctx.Value("sessionId").(string)
	util.PrintJson(location)

	fmt.Printf("\nUpdate body.json to add location. Session ID is %v\n\n", ctx.Value("sessionId"))

	if answer := com.ContinueOrReturn(); answer == 0 {
		return 0
	}

	util.PrintJson(util.ReadJson(&Location{}, "body").(*Location), "Request")
	handle.AddLocation(db)
	response := util.ReadJson(&model.Response{}, "response").(*model.Response)
	util.PrintJson(response, "Response")

	switch response.StatusCode {
	case 200:
		gola.Wait("Press any key to continue")
	case 401:
		gola.Wait(gola.Tf(gola.Color, "Create location failed", gola.LightYellow))
	}
	return -1
}
