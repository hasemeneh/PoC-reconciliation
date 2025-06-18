package bankstatement

import (
	"github.com/jmoiron/sqlx"
)

type bankstatementDomain struct {
	DB *sqlx.DB
}

func New(DB *sqlx.DB) *bankstatementDomain {
	return &bankstatementDomain{
		DB: DB,
	}
}
