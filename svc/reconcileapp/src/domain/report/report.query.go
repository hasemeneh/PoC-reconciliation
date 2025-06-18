package report

import (
	"context"
	"time"

	"github.com/hasemeneh/PoC-OnlineStore/helper/database"
	"github.com/hasemeneh/PoC-reconciliation/svc/reconcileapp/src/models"
	"github.com/jmoiron/sqlx"
)

// SQL Queries
const (
	insertReconciliationReportQuery = `
		INSERT INTO ReconciliationReport (
			upload_date,report_date_start,report_date_end, total_transactions, matched_transactions, unmatched_transactions, total_discrepancies
		) VALUES (?, ?, ?, ?, ?, ?, ?);`

	insertUnmatchedSystemTransactionQuery = `
		INSERT INTO UnmatchedSystemTransaction (
			report_id, trxID, amount, transactionTime, type, bankID
		) VALUES (?, ?, ?, ?, ?, ?);`

	insertUnmatchedBankStatementQuery = `
		INSERT INTO UnmatchedBankStatement (
			report_id, unique_identifier, amount, date, bankID
		) VALUES (?, ?, ?, ?, ?);`

	getReconciliationReportByIDQuery = `
		SELECT id, upload_date, total_transactions, matched_transactions, unmatched_transactions, total_discrepancies
		FROM ReconciliationReport WHERE id = ?;`

	getUnmatchedSystemTransactionsByReportIDQuery = `
		SELECT id, report_id, trxID, amount, transactionTime, type, bankID
		FROM UnmatchedSystemTransaction WHERE report_id = ?;`

	getUnmatchedBankStatementsByReportIDQuery = `
		SELECT id, report_id, unique_identifier, amount, date, bankID
		FROM UnmatchedBankStatement WHERE report_id = ?;`

	updateReconciliationReportQuery = `
		UPDATE ReconciliationReport
		SET upload_date = ?, total_transactions = ?, matched_transactions = ?, unmatched_transactions = ?, total_discrepancies = ?
		WHERE id = ?;`

	updateUnmatchedSystemTransactionQuery = `
		UPDATE UnmatchedSystemTransaction
		SET trxID = ?, amount = ?, transactionTime = ?, type = ?, bankID = ?
		WHERE id = ?;`

	updateUnmatchedBankStatementQuery = `
		UPDATE UnmatchedBankStatement
		SET unique_identifier = ?, amount = ?, date = ?, bankID = ?
		WHERE id = ?;`
)

// Helper function to determine execution context
func getExecFunc(dbtx *sqlx.Tx, db *sqlx.DB) database.ExecFunc {
	if dbtx == nil {
		return db.ExecContext
	}
	return dbtx.ExecContext
}

// Insert Functions
func (p *reportDomain) InsertReconciliationReport(ctx context.Context, dbtx *sqlx.Tx, req *models.ReconciliationReport) (int64, error) {
	execFunc := getExecFunc(dbtx, p.DB)
	result, err := execFunc(ctx, insertReconciliationReportQuery, req.UploadDate, req.ReportDateStart, req.ReportDateEnd, req.TotalTransactions, req.MatchedTransactions, req.UnmatchedTransactions, req.TotalDiscrepancies)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (p *reportDomain) InsertUnmatchedSystemTransaction(ctx context.Context, dbtx *sqlx.Tx, req *models.UnmatchedSystemTransaction) error {
	execFunc := getExecFunc(dbtx, p.DB)
	_, err := execFunc(ctx, insertUnmatchedSystemTransactionQuery, req.ReportID, req.TrxID, req.Amount, req.TransactionTime, req.Type, req.BankID)
	return err
}

func (p *reportDomain) InsertUnmatchedBankStatement(ctx context.Context, dbtx *sqlx.Tx, req *models.UnmatchedBankStatement) error {
	execFunc := getExecFunc(dbtx, p.DB)
	_, err := execFunc(ctx, insertUnmatchedBankStatementQuery, req.ReportID, req.UniqueIdentifier, req.Amount, req.Date, req.BankID)
	return err
}

// Get Functions
func (p *reportDomain) GetReconciliationReportByID(ctx context.Context, id int64) (*models.ReconciliationReport, error) {
	var report models.ReconciliationReport
	err := p.DB.GetContext(ctx, &report, getReconciliationReportByIDQuery, id)
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (p *reportDomain) GetReconciliationReportsByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*models.ReconciliationReport, error) {
	const query = `
		SELECT id, upload_date,report_date_start,report_date_end, total_transactions, matched_transactions, unmatched_transactions, total_discrepancies
		FROM ReconciliationReport
		WHERE (report_date_start > ? AND report_date_start < ?) OR (report_date_end < ? AND report_date_end > ?);`
	var reports []*models.ReconciliationReport
	err := p.DB.SelectContext(ctx, &reports, query, startDate, endDate, endDate, startDate)
	if err != nil {
		return nil, err
	}
	return reports, nil
}

func (p *reportDomain) GetUnmatchedSystemTransactionsByReportID(ctx context.Context, reportID int64) ([]*models.UnmatchedSystemTransaction, error) {
	var transactions []*models.UnmatchedSystemTransaction
	err := p.DB.SelectContext(ctx, &transactions, getUnmatchedSystemTransactionsByReportIDQuery, reportID)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (p *reportDomain) GetUnmatchedBankStatementsByReportID(ctx context.Context, reportID int64) ([]*models.UnmatchedBankStatement, error) {
	var statements []*models.UnmatchedBankStatement
	err := p.DB.SelectContext(ctx, &statements, getUnmatchedBankStatementsByReportIDQuery, reportID)
	if err != nil {
		return nil, err
	}
	return statements, nil
}

func (p *reportDomain) GetAllUnmatchedByReportID(ctx context.Context, reportID int64) (*models.UnmatchedData, error) {
	systemTransactions, err := p.GetUnmatchedSystemTransactionsByReportID(ctx, reportID)
	if err != nil {
		return nil, err
	}

	bankStatements, err := p.GetUnmatchedBankStatementsByReportID(ctx, reportID)
	if err != nil {
		return nil, err
	}

	return &models.UnmatchedData{
		UnmatchedSystemTransactions: systemTransactions,
		UnmatchedBankStatements:     bankStatements,
	}, nil
}

// Update Functions
func (p *reportDomain) UpdateReconciliationReport(ctx context.Context, dbtx *sqlx.Tx, req *models.ReconciliationReport) error {
	execFunc := getExecFunc(dbtx, p.DB)
	_, err := execFunc(ctx, updateReconciliationReportQuery, req.UploadDate, req.TotalTransactions, req.MatchedTransactions, req.UnmatchedTransactions, req.TotalDiscrepancies, req.ID)
	return err
}

func (p *reportDomain) UpdateUnmatchedSystemTransaction(ctx context.Context, dbtx *sqlx.Tx, req *models.UnmatchedSystemTransaction) error {
	execFunc := getExecFunc(dbtx, p.DB)
	_, err := execFunc(ctx, updateUnmatchedSystemTransactionQuery, req.TrxID, req.Amount, req.TransactionTime, req.Type, req.BankID, req.ID)
	return err
}

func (p *reportDomain) UpdateUnmatchedBankStatement(ctx context.Context, dbtx *sqlx.Tx, req *models.UnmatchedBankStatement) error {
	execFunc := getExecFunc(dbtx, p.DB)
	_, err := execFunc(ctx, updateUnmatchedBankStatementQuery, req.UniqueIdentifier, req.Amount, req.Date, req.BankID, req.ID)
	return err
}

// Transaction Management
func (p *reportDomain) StartTx(ctx context.Context) (*sqlx.Tx, error) {
	return p.DB.BeginTxx(ctx, nil)
}

func (p *reportDomain) GetUnmatchedSystemTransactionsByReportIDAndDateRange(ctx context.Context, reportID int64, startDate, endDate time.Time) ([]*models.UnmatchedSystemTransaction, error) {
	const query = `
		SELECT id, report_id, trxID, amount, transactionTime, type, bankID
		FROM UnmatchedSystemTransaction
		WHERE report_id = ? AND transactionTime BETWEEN ? AND ?;`
	var transactions []*models.UnmatchedSystemTransaction
	err := p.DB.SelectContext(ctx, &transactions, query, reportID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (p *reportDomain) GetUnmatchedBankStatementsByReportIDAndDateRange(ctx context.Context, reportID int64, startDate, endDate time.Time) ([]*models.UnmatchedBankStatement, error) {
	const query = `
		SELECT id, report_id, unique_identifier, amount, date, bankID
		FROM UnmatchedBankStatement
		WHERE report_id = ? AND date BETWEEN ? AND ?;`
	var statements []*models.UnmatchedBankStatement
	err := p.DB.SelectContext(ctx, &statements, query, reportID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return statements, nil
}
