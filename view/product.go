package view

import (
	"context"
	"database/sql"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
	"project/util"
)

type Product struct {
}

func (screen Product) Render(ctx context.Context, db *sql.DB) int {
	util.H1("Manajemen Stok Barang")

	fmt.Println("[1] Tambah Barang")
	fmt.Println("[2] Update Stok")
	fmt.Println("[3] Penerimaan Barang")
	fmt.Println("[4] Pengeluaran Barang")

	fmt.Println("\n[0] Kembali")

	menuItem, _ := gola.ToInt(gola.Input(gola.Args(gola.P("type", "number"), gola.P("label", fmt.Sprintf("%s :", "Pilih [1]-[4] atau [0] untuk kembali")))))

	switch menuItem {
	case 0:
		return 0
	case 1:
	case 2:
	case 3:
	case 4:
	}
	return -1
}
