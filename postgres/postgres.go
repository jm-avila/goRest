package postgres

import (
	"github.com/go-pg/pg/v9"
	// not sure why i need it
	_ "github.com/lib/pq"
)

// New db connection
func New(opts *pg.Options) *pg.DB {
	db := pg.Connect(opts)

	return db
}
