package view

import (
	"context"
	"database/sql"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
	handle "project/handler/products"
	"project/model"
	"project/util"
	com "project/util"
)

type ProductSearch struct {
	Name      string `json:"name"`
	SessionId string `json:"session_id"`
}

func (product ProductSearch) Render(ctx context.Context, db *sql.DB) int {
	com.H1("Pencarian Barang")

	util.PrintJson(model.Search{Name: "celana pendek", Page: 1, Limit: 5, SessionId: ctx.Value("sessionId").(string)})

	fmt.Printf("\nUpdate body.json to search for product by name. Session ID is %v\n\n", ctx.Value("sessionId"))

	if answer := com.ContinueOrReturn(); answer == 0 {
		return 0
	}
	util.PrintJson(util.ReadJson(&model.Search{}, "body").(*model.Search), "Request")
	handle.ProductSearch(db)
	response := util.ReadJson(&model.DataPage{}, "response").(*model.DataPage)
	util.PrintJson(response, "Response")
	switch response.StatusCode {
	case 200:
		gola.Wait("Press any key to continue")
	case 401:
		gola.Wait(gola.Tf(gola.Color, "Create category failed", gola.LightYellow))
	}
	return -1
}
