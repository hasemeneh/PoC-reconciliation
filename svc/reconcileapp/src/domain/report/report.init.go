package report

import (
	"github.com/jmoiron/sqlx"
)

type reportDomain struct {
	DB *sqlx.DB
}

func New(DB *sqlx.DB) *reportDomain {
	return &reportDomain{
		DB: DB,
	}
}
