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

type TransactionIndex struct {
}

func (transaction TransactionIndex) Render(ctx context.Context, db *sql.DB) int {
	com.H1("Riwayat Transaksi")

	util.PrintJson(model.Search{Page: 1, Limit: 5, SessionId: ctx.Value("sessionId").(string)})

	fmt.Printf("\nUpdate body.json to get transaction history. Session ID is %v\n\n", ctx.Value("sessionId"))

	if answer := com.ContinueOrReturn(); answer == 0 {
		return 0
	}

	util.PrintJson(util.ReadJson(&model.Search{}, "body").(*model.Search), "Request")
	handle.TransactionIndex(db)
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
