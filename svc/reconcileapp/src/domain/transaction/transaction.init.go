package transaction

import (
	"github.com/jmoiron/sqlx"
)

type transactionDomain struct {
	DB *sqlx.DB
}

func New(DB *sqlx.DB) *transactionDomain {
	return &transactionDomain{
		DB: DB,
	}
}
