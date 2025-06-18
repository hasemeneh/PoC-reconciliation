package models

type Bank struct {
	BankID int    `db:"bankID" json:"bank_id"`
	Name   string `db:"name" json:"name"`
	Code   string `db:"code" json:"code"`
}
