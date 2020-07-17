package postgres

import "github.com/go-pg/pg"

type UserRepo struct {
	DB *pg.DB
}

func NewUserRepo(DB *pg.DB) *UserRepo {
	return &UserRepo{DB: DB}
}
