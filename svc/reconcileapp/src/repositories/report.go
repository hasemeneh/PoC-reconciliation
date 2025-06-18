package repositories

import (
	"context"
	"time"

	"github.com/hasemeneh/PoC-reconciliation/svc/reconcileapp/src/models"
	"github.com/jmoiron/sqlx"
)

// ReportRepositories defines the methods for the report domain
type ReportRepositories interface {
	InsertReconciliationReport(ctx context.Context, dbtx *sqlx.Tx, req *models.ReconciliationReport) (int64, error)
	InsertUnmatchedSystemTransaction(ctx context.Context, dbtx *sqlx.Tx, req *models.UnmatchedSystemTransaction) error
	InsertUnmatchedBankStatement(ctx context.Context, dbtx *sqlx.Tx, req *models.UnmatchedBankStatement) error

	GetReconciliationReportByID(ctx context.Context, id int64) (*models.ReconciliationReport, error)
	GetUnmatchedSystemTransactionsByReportID(ctx context.Context, reportID int64) ([]*models.UnmatchedSystemTransaction, error)
	GetUnmatchedBankStatementsByReportID(ctx context.Context, reportID int64) ([]*models.UnmatchedBankStatement, error)
	GetAllUnmatchedByReportID(ctx context.Context, reportID int64) (*models.UnmatchedData, error)
	GetReconciliationReportsByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*models.ReconciliationReport, error)

	UpdateReconciliationReport(ctx context.Context, dbtx *sqlx.Tx, req *models.ReconciliationReport) error
	UpdateUnmatchedSystemTransaction(ctx context.Context, dbtx *sqlx.Tx, req *models.UnmatchedSystemTransaction) error
	UpdateUnmatchedBankStatement(ctx context.Context, dbtx *sqlx.Tx, req *models.UnmatchedBankStatement) error

	GetUnmatchedSystemTransactionsByReportIDAndDateRange(ctx context.Context, reportID int64, startDate, endDate time.Time) ([]*models.UnmatchedSystemTransaction, error)
	GetUnmatchedBankStatementsByReportIDAndDateRange(ctx context.Context, reportID int64, startDate, endDate time.Time) ([]*models.UnmatchedBankStatement, error)

	StartTx(ctx context.Context) (*sqlx.Tx, error)
}
