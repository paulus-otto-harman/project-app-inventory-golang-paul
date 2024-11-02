package view

import (
	"context"
	"database/sql"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
	"project/util"
	viewcategory "project/view/category"
	viewlocation "project/view/location"
)

type Category struct {
}

func (screen Category) Render(ctx context.Context, db *sql.DB) int {
	util.H1("Manajemen Kategori dan Lokasi Barang")

	fmt.Println("[1] Kategori Barang")
	fmt.Println("[2] Lokasi Penyimpanan")

	fmt.Println("\n[0] Kembali")

	menuItem, _ := gola.ToInt(gola.Input(gola.Args(gola.P("type", "number"), gola.P("label", fmt.Sprintf("%s :", "Pilih [1]-[2] atau [0] untuk kembali")))))

	switch menuItem {
	case 0:
		return 0
	case 1:
		Render(&viewcategory.Category{}, ctx, db)
	case 2:
		Render(&viewlocation.Location{}, ctx, db)
	}
	return -1
}
