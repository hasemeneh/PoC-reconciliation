package bank

import (
	"github.com/jmoiron/sqlx"
)

type bankDomain struct {
	DB *sqlx.DB
}

func New(DB *sqlx.DB) *bankDomain {
	return &bankDomain{
		DB: DB,
	}
}
