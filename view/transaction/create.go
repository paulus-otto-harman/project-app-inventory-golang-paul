package transaction

import (
	"context"
	"database/sql"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
	handle "project/handler/transactions"
	"project/model"
	"project/util"
	com "project/util"
)

type Transaction struct {
}

func (transaction Transaction) Render(ctx context.Context, db *sql.DB) int {
	com.H1("Pencatatan Keluar/Masuk Barang")

	util.PrintJson(struct {
		Type      string `json:"type"`
		ProductId int    `json:"product_id"`
		Quantity  int    `json:"quantity"`
		Note      string `json:"note"`
		SessionId string `json:"session_id"`
	}{
		Type:      "out",
		SessionId: ctx.Value("sessionId").(string),
	})

	fmt.Printf("\nUpdate body.json to create transaction. Session ID is %v\n\n", ctx.Value("sessionId"))

	if answer := com.ContinueOrReturn(); answer == 0 {
		return 0
	}

	util.PrintJson(util.ReadJson(&model.Transaction{}, "body").(*model.Transaction), "Request")
	handle.TransactionCreate(db)
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
