package definitions

import (
	"context"
	"time"

	"github.com/hasemeneh/PoC-reconciliation/svc/reconcileapp/src/models"
)

type ReconcileDefinition interface {
	ReconcileTransactions(transactionsRecord models.CSVData, bankStatement []*models.CSVData, startDate time.Time, endDate time.Time) error
	GetUnmatchedReportByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*models.ReconciliationReportResponse, error)
}
