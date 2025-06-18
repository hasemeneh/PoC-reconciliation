package repositories

import (
	"context"

	"github.com/hasemeneh/PoC-reconciliation/svc/reconcileapp/src/models"
	"github.com/jmoiron/sqlx"
)

type TransactionsRepo interface {
	GetTransactionByID(ctx context.Context, trxID string) (*models.TransactionModel, error)
	GetAllTransactions(ctx context.Context) ([]*models.TransactionModel, error)
	UpdateTransaction(ctx context.Context, dbtx *sqlx.Tx, req *models.TransactionModel) error
	DeleteTransaction(ctx context.Context, dbtx *sqlx.Tx, trxID string) error
	InsertTransaction(ctx context.Context, dbtx *sqlx.Tx, req *models.TransactionModel) error
	StartTx(ctx context.Context) (*sqlx.Tx, error)
}
