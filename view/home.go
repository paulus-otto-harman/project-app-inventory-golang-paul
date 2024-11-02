package view

import (
	"context"
	"database/sql"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
	"project/util"
)

type Home struct {
	isLogout bool
}

func (home *Home) Render(ctx context.Context, db *sql.DB) int {
	util.H1("Menu Utama")

	fmt.Println("[1] Manajemen Stok Barang")
	fmt.Println("[2] Manajemen Kategori dan Lokasi Barang")
	fmt.Println("[3] Pencarian dan Riwayat Transaksi")

	fmt.Println("\n[0] Logout")

	menuItem, _ := gola.ToInt(gola.Input(gola.Args(gola.P("type", "number"), gola.P("label", fmt.Sprintf("%s :", "Pilih [1]-[3] atau [0] untuk Logout")))))
	switch menuItem {
	case 0:
		logout := Logout{}
		Render(&logout, ctx, db)
		if logout.Confirmed {
			home.isLogout = true
			return 0
		}
	case 1:
		Render(&Product{}, ctx, db)
	case 2:
		Render(&Category{}, ctx, db)
	case 3:
	case 9:
		home.isLogout = true
		return 0
	}
	return -1
}
