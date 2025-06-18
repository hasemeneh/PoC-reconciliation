package models

import (
	"time"
)

// BankStatement represents the BankStatement table in the database.
type BankStatement struct {
	UniqueIdentifier string    `json:"unique_identifier" db:"unique_identifier"`
	Amount           float64   `json:"amount" db:"amount"`
	Date             time.Time `json:"statement_date" db:"statement_date"`
	IsReconciled     bool      `json:"is_reconciled" db:"isReconciled"`
	BankID           int       `json:"bank_id" db:"bankID"`
}

func (c CSVData) ToBankStatementModels() ([]*BankStatement, error) {
	var bankStatements []*BankStatement
	for k, row := range c {
		if k == 0 {
			continue // Skip header row
		}

		if len(row) < 4 {
			continue // Skip rows with insufficient data
		}
		// Parse the bank statement model from the CSV row
		amount, err := parseAmount(row[1])
		if err != nil {
			return nil, err // Return error if parsing fails
		}
		date, err := parseTime(row[2])
		if err != nil {
			return nil, err // Return error if parsing fails
		}
		bankID, err := parseBankID(row[3])
		if err != nil {
			return nil, err // Return error if parsing fails
		}
		bankStatement := &BankStatement{
			UniqueIdentifier: row[0],
			Amount:           amount,
			Date:             date,
			IsReconciled:     false,
			BankID:           bankID,
		}
		bankStatements = append(bankStatements, bankStatement)
	}
	return bankStatements, nil
}
