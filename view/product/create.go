package view

import (
	"context"
	"database/sql"
	"fmt"
	handle "project/handler/products"
	"project/model"
	"project/util"
	com "project/util"
)

type Product struct {
	Id        int            `json:"id"`
	Name      string         `json:"name"`
	Category  model.Category `json:"category"`
	Location  model.Location `json:"location"`
	SessionId string         `json:"session_id"`
}

func (product Product) Render(ctx context.Context, db *sql.DB) int {
	com.H1("Lokasi Penyimpanan")

	product.SessionId = ctx.Value("sessionId").(string)
	util.PrintJson(product)

	fmt.Printf("\nUpdate body.json to add product. Session ID is %v\n\n", ctx.Value("sessionId"))

	if answer := com.ContinueOrReturn(); answer == 0 {
		return 0
	}

	util.PrintJson(util.ReadJson(&Product{}, "body").(*Product), "Request")
	handle.AddProduct(db)
	response := util.ReadJson(&model.Response{}, "response").(*model.Response)
	util.PrintJson(response, "Response")

	return -1
}
