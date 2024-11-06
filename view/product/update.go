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

type ProductUpdate struct {
	Id        int    `json:"id"`
	Balance   int    `json:"balance"`
	SessionId string `json:"session_id"`
}

func (product ProductUpdate) Render(ctx context.Context, db *sql.DB) int {
	com.H1("Update Stok")

	product.SessionId = ctx.Value("sessionId").(string)
	util.PrintJson(product)

	fmt.Printf("\nUpdate body.json to update balance of a product. Session ID is %v\n\n", ctx.Value("sessionId"))

	if answer := com.ContinueOrReturn(); answer == 0 {
		return 0
	}

	util.PrintJson(util.ReadJson(&ProductUpdate{}, "body").(*ProductUpdate), "Request")
	handle.ProductUpdate(db)
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
