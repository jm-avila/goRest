package postgres

import (
	"github.com/go-pg/pg"
)

// New db connection
func New(opts *pg.Options) *pg.DB {
	db := pg.Connect(opts)

	return db
}
