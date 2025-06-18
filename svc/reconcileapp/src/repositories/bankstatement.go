package repositories

import (
	"context"

	"github.com/hasemeneh/PoC-reconciliation/svc/reconcileapp/src/models"
	"github.com/jmoiron/sqlx"
)

type BankStatementRepositories interface {
	GetStatementByID(ctx context.Context, uniqueIdentifier string) (*models.BankStatement, error)
	GetAllStatements(ctx context.Context) ([]*models.BankStatement, error)
	UpdateStatement(ctx context.Context, dbtx *sqlx.Tx, req *models.BankStatement) error
	DeleteStatement(ctx context.Context, dbtx *sqlx.Tx, uniqueIdentifier string) error
	InsertStatement(ctx context.Context, dbtx *sqlx.Tx, req *models.BankStatement) error
	StartTx(ctx context.Context) (*sqlx.Tx, error)
}
