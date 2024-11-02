package view

import (
	"context"
	"database/sql"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
	handle "project/handler/categories"
	"project/model"
	"project/util"
	com "project/util"
)

type Category struct {
	Id        int    `json:"-"`
	SessionId string `json:"session_id"`
	Name      string `json:"name"`
}

func (category Category) Render(ctx context.Context, db *sql.DB) int {
	com.H1("Kategori Barang")
	category.SessionId = ctx.Value("sessionId").(string)
	util.PrintJson(category)

	fmt.Printf("\nUpdate body.json to add category. Session ID is %v\n\n", ctx.Value("sessionId"))

	if answer := com.ContinueOrReturn(); answer == 0 {
		return 0
	}

	util.PrintJson(util.ReadJson(&Category{}, "body").(*Category), "Request")
	handle.AddCategory(db)
	response := util.ReadJson(&model.Response{}, "response").(*model.Response)
	util.PrintJson(response, "Response")

	switch response.StatusCode {
	case 200:
		gola.Wait("Press any key to continue")
	case 401:
		gola.Wait(gola.Tf(gola.Color, "Create category failed", gola.LightYellow))
	}
	return -1
}
