package repositories

import (
	"context"

	"github.com/hasemeneh/PoC-reconciliation/svc/reconcileapp/src/models"
	"github.com/jmoiron/sqlx"
)

// BankRepositories defines the interface for bankDomain methods
type BankRepositories interface {
	GetBankByID(ctx context.Context, bankID int) (*models.Bank, error)
	GetAllBanks(ctx context.Context) ([]*models.Bank, error)
	InsertBank(ctx context.Context, dbtx *sqlx.Tx, req *models.Bank) error
	UpdateBank(ctx context.Context, dbtx *sqlx.Tx, req *models.Bank) error
	DeleteBank(ctx context.Context, dbtx *sqlx.Tx, bankID int) error
}
