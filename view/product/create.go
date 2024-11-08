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

type Product struct {
	Id         int    `json:"-"`
	Name       string `json:"name"`
	CategoryId int    `json:"category_id"`
	LocationId int    `json:"location_id"`
	SessionId  string `json:"session_id"`
}

func (product Product) Render(ctx context.Context, db *sql.DB) int {
	com.H1("Tambah Barang")

	product.SessionId = ctx.Value("sessionId").(string)
	util.PrintJson(product)

	fmt.Printf("\nUpdate body.json to add product. Session ID is %v\n\n", ctx.Value("sessionId"))

	if answer := com.ContinueOrReturn(); answer == 0 {
		return 0
	}

	util.PrintJson(util.ReadJson(&Product{}, "body").(*Product), "Request")
	handle.ProductCreate(db)
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
