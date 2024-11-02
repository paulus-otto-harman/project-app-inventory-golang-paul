package view

import (
	"context"
	"database/sql"
)

type Screen interface {
	Render(ctx context.Context, db *sql.DB) int
}

func Render(screen Screen, ctx context.Context, db *sql.DB) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if screen.Render(ctx, db) == 0 {
				return
			}
		}
	}
}
