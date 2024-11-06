package view

import (
	"context"
	"database/sql"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
	"project/util"
	viewproduct "project/view/product"
	"project/view/transaction"
)

type Product struct {
}

func (screen Product) Render(ctx context.Context, db *sql.DB) int {
	util.H1("Manajemen Stok Barang")

	fmt.Println("[1] Tambah Barang")
	fmt.Println("[2] Update Stok")
	fmt.Println("[3] Pencatatan Keluar/Masuk Barang")

	fmt.Println("\n[0] Kembali")

	menuItem, _ := gola.ToInt(gola.Input(gola.Args(gola.P("type", "number"), gola.P("label", fmt.Sprintf("%s :", "Pilih [1]-[3] atau [0] untuk kembali")))))

	switch menuItem {
	case 0:
		return 0
	case 1:
		Render(&viewproduct.Product{}, ctx, db)
	case 2:
		Render(&viewproduct.ProductUpdate{}, ctx, db)
	case 3:
		Render(&transaction.Transaction{}, ctx, db)
	}
	return -1
}
