package view

import (
	"context"
	"database/sql"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
	"project/util"
	"project/view/product"
	viewtransaction "project/view/transaction"
)

type Search struct {
}

func (screen Search) Render(ctx context.Context, db *sql.DB) int {
	util.H1("Pencarian dan Riwayat Transaksi")

	fmt.Println("[1] Pencarian Barang")
	fmt.Println("[2] Riwayat Transaksi")

	fmt.Println("\n[0] Kembali")

	menuItem, _ := gola.ToInt(gola.Input(gola.Args(gola.P("type", "number"), gola.P("label", fmt.Sprintf("%s :", "Pilih [1]-[2] atau [0] untuk kembali")))))

	switch menuItem {
	case 0:
		return 0
	case 1:
		Render(&view.ProductSearch{}, ctx, db)
	case 2:
		Render(&viewtransaction.TransactionIndex{}, ctx, db)
	}
	return -1
}
