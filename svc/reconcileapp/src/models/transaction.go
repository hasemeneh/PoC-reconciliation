package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type TransactionModel struct {
	TrxID           string    `json:"trx_id" db:"trxID"`                     // Primary key
	Amount          float64   `json:"amount" db:"amount"`                    // DECIMAL(15, 2)
	Type            string    `json:"type" db:"type"`                        // ENUM('DEBIT', 'CREDIT')
	TransactionTime time.Time `json:"transaction_time" db:"transactionTime"` // DATETIME
	IsReconciled    bool      `json:"is_reconciled" db:"isReconciled"`       // BOOLEAN, default false
	BankID          int       `json:"bank_id" db:"bankID"`                   // Foreign key reference to Bank(bankID)
}

const (
	TransactionTypeDebit  = "DEBIT"
	TransactionTypeCredit = "CREDIT"
)

func (c CSVData) ToTransactionModels() ([]*TransactionModel, error) {
	var transactions []*TransactionModel
	for k, row := range c {
		if k == 0 {
			continue // Skip header row
		}
		if len(row) < 5 {
			continue // Skip rows with insufficient data
		}
		// Parse the transaction model from the CSV row
		amount, err := parseAmount(row[1])
		if err != nil {
			return nil, err // Return error if amount parsing fails
		}

		transactionTime, err := parseTime(row[3])
		if err != nil {
			return nil, err // Return error if transaction time parsing fails
		}

		bankID, err := parseBankID(row[4])
		if err != nil {
			return nil, err // Return error if bank ID parsing fails
		}

		// Validate transaction type
		transactionType := strings.ToUpper(strings.ReplaceAll(strings.TrimSpace(row[2]), "'", ""))
		if transactionType != "DEBIT" && transactionType != "CREDIT" {
			return nil, fmt.Errorf("invalid transaction type: %s", transactionType)
		}

		transaction := &TransactionModel{
			TrxID:           row[0],
			Amount:          amount,
			Type:            transactionType,
			TransactionTime: transactionTime,
			BankID:          bankID,
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func parseAmount(amountStr string) (float64, error) {
	amountStr = strings.TrimSpace(amountStr)
	amountStr = strings.ReplaceAll(amountStr, ",", "")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return 0, err // Return 0 if parsing fails
	}
	return amount, nil
}

func parseTime(timeStr string) (time.Time, error) {
	timeStr = strings.TrimSpace(timeStr)
	timeStr = strings.ReplaceAll(timeStr, "'", "")

	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr) // Assuming the format is "YYYY-MM-DD HH:MM:SS"
	if err != nil {
		return time.Time{}, err // Return zero value if parsing fails
	}
	return parsedTime, nil
}

func parseBankID(bankIDStr string) (int, error) {
	bankIDStr = strings.TrimSpace(bankIDStr)

	bankID, err := strconv.Atoi(bankIDStr)
	if err != nil {
		return 0, err // Return 0 if parsing fails
	}
	return bankID, nil
}
