package models

import "time"

// ReconciliationReport represents the main reconciliation report table
type ReconciliationReport struct {
	ID                    int64     `json:"id" db:"id"`
	UploadDate            time.Time `json:"upload_date" db:"upload_date"`
	ReportDateStart       time.Time `json:"report_date_start" db:"report_date_start"`
	ReportDateEnd         time.Time `json:"report_date_end" db:"report_date_end"`
	TotalTransactions     int       `json:"total_transactions" db:"total_transactions"`
	MatchedTransactions   int       `json:"matched_transactions" db:"matched_transactions"`
	UnmatchedTransactions int       `json:"unmatched_transactions" db:"unmatched_transactions"`
	TotalDiscrepancies    float64   `json:"total_discrepancies" db:"total_discrepancies"`
}

// UnmatchedSystemTransaction represents the unmatched system transactions table
type UnmatchedSystemTransaction struct {
	ID              int       `json:"id" db:"id"`
	ReportID        int       `json:"report_id" db:"report_id"`
	TrxID           string    `json:"trx_id" db:"trxID"`
	Amount          float64   `json:"amount" db:"amount"`
	TransactionTime time.Time `json:"transaction_time" db:"transactionTime"`
	Type            string    `json:"type" db:"type"` // DEBIT or CREDIT
	BankID          int       `json:"bank_id" db:"bankID"`
}

// UnmatchedBankStatement represents the unmatched bank statement records table
type UnmatchedBankStatement struct {
	ID               int       `json:"id" db:"id"`
	ReportID         int       `json:"report_id" db:"report_id"`
	UniqueIdentifier string    `json:"unique_identifier" db:"unique_identifier"`
	Amount           float64   `json:"amount" db:"amount"`
	Date             time.Time `json:"date" db:"date"`
	BankID           int       `json:"bank_id" db:"bankID"`
}

// ConvertBankStatementToUnmatched converts a single BankStatement to an UnmatchedBankStatement
func ConvertBankStatementToUnmatched(reportID int, statement *BankStatement) *UnmatchedBankStatement {
	return &UnmatchedBankStatement{
		ReportID:         reportID,
		UniqueIdentifier: statement.UniqueIdentifier,
		Amount:           statement.Amount,
		Date:             statement.Date,
		BankID:           statement.BankID,
	}
}

// ConvertTransactionToUnmatched converts a single TransactionModel to an UnmatchedSystemTransaction
func ConvertTransactionToUnmatched(reportID int, transaction *TransactionModel) *UnmatchedSystemTransaction {
	return &UnmatchedSystemTransaction{
		ReportID:        reportID,
		TrxID:           transaction.TrxID,
		Amount:          transaction.Amount,
		TransactionTime: transaction.TransactionTime,
		Type:            transaction.Type,
		BankID:          transaction.BankID,
	}
}

// UnmatchedData represents the combined unmatched data containing lists of unmatched system transactions and bank statements
type UnmatchedData struct {
	UnmatchedSystemTransactions []*UnmatchedSystemTransaction `json:"unmatched_system_transactions"`
	UnmatchedBankStatements     []*UnmatchedBankStatement     `json:"unmatched_bank_statements"`
}
type CSVData [][]string

// NewCSVData converts a [][]string to CSVData
func NewCSVData(data [][]string) CSVData {
	return CSVData(data)
}

type ReconciliationReportPrepared struct {
	Transactions   []*TransactionModel
	BankStatements []*BankStatement
	BankID         int
	Amount         float64
	Date           time.Time
	DateString     string
}

type ReconciliationReportResponse struct {
	Transactions   []*UnmatchedSystemTransaction
	BankStatements []*UnmatchedBankStatement
	ReportID       int64
	Amount         float64
	Date           time.Time
}
