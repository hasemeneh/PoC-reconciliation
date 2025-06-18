package reconciliation

import (
	"context"
	"fmt"
	"time"

	"github.com/hasemeneh/PoC-reconciliation/svc/reconcileapp/src/models"
)

// ReconcileTransactions processes the reconciliation between system transactions and bank statements.
func (r *reconciliationUsecase) ReconcileTransactions(transactionsRecord models.CSVData, bankStatement []*models.CSVData, startDate time.Time, endDate time.Time) error {
	ctx := context.Background()
	transactions, err := transactionsRecord.ToTransactionModels()
	if err != nil {
		return err
	}
	var bankStatements []*models.BankStatement
	for _, statement := range bankStatement {
		bankStatementModel, err := statement.ToBankStatementModels()
		if err != nil {
			return err
		}
		bankStatements = append(bankStatements, bankStatementModel...)
	}

	// Use a single map to store reconciliation data
	reconciliationMap := make(map[string]*models.ReconciliationReportPrepared)

	// Process transactions
	for _, transaction := range transactions {
		amount := transaction.Amount
		if transaction.Type == models.TransactionTypeCredit {
			amount = -amount // Convert credit transactions to negative amounts
		}
		key := generateKey(transaction.BankID, transaction.TransactionTime, amount)
		if _, exists := reconciliationMap[key]; !exists {
			reconciliationMap[key] = &models.ReconciliationReportPrepared{
				BankID:         transaction.BankID,
				Amount:         transaction.Amount,
				Date:           transaction.TransactionTime,
				DateString:     transaction.TransactionTime.Format("2006-01-02"),
				Transactions:   nil,
				BankStatements: nil,
			}
		}
		reconciliationMap[key].Transactions = append(reconciliationMap[key].Transactions, transaction)
	}

	// Process bank statements
	for _, statement := range bankStatements {
		key := generateKey(statement.BankID, statement.Date, statement.Amount)
		if _, exists := reconciliationMap[key]; !exists {
			reconciliationMap[key] = &models.ReconciliationReportPrepared{
				BankID:         statement.BankID,
				Amount:         statement.Amount,
				Date:           statement.Date,
				DateString:     statement.Date.Format("2006-01-02"),
				Transactions:   nil,
				BankStatements: nil,
			}
		}
		reconciliationMap[key].BankStatements = append(reconciliationMap[key].BankStatements, statement)
	}

	tx, err := r.Transactions.StartTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	reportObj := &models.ReconciliationReport{
		UploadDate:      time.Now(),
		ReportDateStart: startDate,
		ReportDateEnd:   endDate,
	}

	reportID, err := r.Report.InsertReconciliationReport(ctx, tx, reportObj)
	if err != nil {
		return err
	}

	// Process reconciliation reports
	for _, report := range reconciliationMap {

		minLen := len(report.Transactions)
		if len(report.BankStatements) < minLen {
			minLen = len(report.BankStatements)
		}
		// Mark matched pairs as reconciled
		for i := 0; i < minLen; i++ {
			report.Transactions[i].IsReconciled = true
			report.BankStatements[i].IsReconciled = true
		}
		diff := len(report.Transactions) - len(report.BankStatements)
		for _, transaction := range report.Transactions {
			err = r.Transactions.InsertTransaction(ctx, tx, transaction)
			if err != nil {
				return err
			}

			if !transaction.IsReconciled {
				err = r.Report.InsertUnmatchedSystemTransaction(ctx, tx, models.ConvertTransactionToUnmatched(int(reportID), transaction))
				if err != nil {
					return err
				}
			}
		}
		for _, statement := range report.BankStatements {
			err = r.BankStatement.InsertStatement(ctx, tx, statement)
			if err != nil {
				return err
			}

			if !statement.IsReconciled {
				err = r.Report.InsertUnmatchedBankStatement(ctx, tx, models.ConvertBankStatementToUnmatched(int(reportID), statement))
				if err != nil {
					return err
				}
			}
		}
		reportObj.TotalTransactions = reportObj.TotalTransactions + len(report.Transactions)
		reportObj.UnmatchedTransactions = reportObj.UnmatchedTransactions + abs(diff)
		reportObj.MatchedTransactions = reportObj.MatchedTransactions + len(report.Transactions) - abs(diff)
		reportObj.TotalDiscrepancies = reportObj.TotalDiscrepancies + float64(abs(diff))*report.Amount

	}

	return tx.Commit()
}

// generateKey creates a unique key for reconciliation data based on bank ID, date, and amount.
func generateKey(bankID int, date time.Time, amount float64) string {
	return fmt.Sprintf("%d|%s|%f", bankID, date.Format("2006-01-02"), amount)
}

// abs returns the absolute value of an integer.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (r *reconciliationUsecase) GetUnmatchedReportByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*models.ReconciliationReportResponse, error) {
	resp := make([]*models.ReconciliationReportResponse, 0)

	reports, err := r.Report.GetReconciliationReportsByDateRange(ctx, startDate, endDate)
	if err != nil {
		return resp, err
	}

	for _, report := range reports {
		unmatchedBankStatements, err := r.Report.GetUnmatchedBankStatementsByReportIDAndDateRange(ctx, report.ID, startDate, endDate)
		if err != nil {
			return resp, err
		}

		unmatchedTransactions, err := r.Report.GetUnmatchedSystemTransactionsByReportIDAndDateRange(ctx, report.ID, startDate, endDate)
		if err != nil {
			return resp, err
		}
		resp = append(resp, &models.ReconciliationReportResponse{
			ReportID:       report.ID,
			Amount:         report.TotalDiscrepancies,
			Date:           report.UploadDate,
			Transactions:   unmatchedTransactions,
			BankStatements: unmatchedBankStatements,
		})

	}

	return resp, nil
}
