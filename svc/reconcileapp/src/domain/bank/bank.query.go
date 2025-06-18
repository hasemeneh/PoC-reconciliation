package bank

import (
	"context"

	"github.com/hasemeneh/PoC-OnlineStore/helper/database"
	"github.com/hasemeneh/PoC-reconciliation/svc/reconcileapp/src/models"
	"github.com/jmoiron/sqlx"
)

// Queries for the Bank table
var getBankByIDQuery = "SELECT `bankID`, `name`, `code` FROM `Bank` WHERE `bankID` = ?;"
var insertBankQuery = "INSERT INTO `Bank` (`name`, `code`) VALUES (?, ?);"
var getAllBanksQuery = "SELECT `bankID`, `name`, `code` FROM `Bank`;"
var updateBankQuery = "UPDATE `Bank` SET `name` = ?, `code` = ? WHERE `bankID` = ?;"
var deleteBankQuery = "DELETE FROM `Bank` WHERE `bankID` = ?;"

// Method to get a bank by its ID
func (p *bankDomain) GetBankByID(ctx context.Context, bankID int) (*models.Bank, error) {
	var bank models.Bank
	err := p.DB.GetContext(ctx, &bank, getBankByIDQuery, bankID)
	if err != nil {
		return nil, err
	}
	return &bank, nil
}

// Method to get all banks
func (p *bankDomain) GetAllBanks(ctx context.Context) ([]*models.Bank, error) {
	var banks []*models.Bank
	err := p.DB.SelectContext(ctx, &banks, getAllBanksQuery)
	if err != nil {
		return nil, err
	}
	return banks, nil
}

// Method to insert a new bank
func (p *bankDomain) InsertBank(ctx context.Context, dbtx *sqlx.Tx, req *models.Bank) error {
	var execFunc database.ExecFunc
	if dbtx == nil {
		execFunc = p.DB.ExecContext
	} else {
		execFunc = dbtx.ExecContext
	}
	_, err := execFunc(ctx, insertBankQuery, req.Name, req.Code)
	return err
}

// Method to update an existing bank
func (p *bankDomain) UpdateBank(ctx context.Context, dbtx *sqlx.Tx, req *models.Bank) error {
	var execFunc database.ExecFunc
	if dbtx == nil {
		execFunc = p.DB.ExecContext
	} else {
		execFunc = dbtx.ExecContext
	}
	_, err := execFunc(ctx, updateBankQuery, req.Name, req.Code, req.BankID)
	return err
}

// Method to delete a bank by its ID
func (p *bankDomain) DeleteBank(ctx context.Context, dbtx *sqlx.Tx, bankID int) error {
	var execFunc database.ExecFunc
	if dbtx == nil {
		execFunc = p.DB.ExecContext
	} else {
		execFunc = dbtx.ExecContext
	}
	_, err := execFunc(ctx, deleteBankQuery, bankID)
	return err
}

// BankRepositories defines the interface for bankDomain methods
type BankRepositories interface {
	GetBankByID(ctx context.Context, bankID int) (*models.Bank, error)
	GetAllBanks(ctx context.Context) ([]*models.Bank, error)
	InsertBank(ctx context.Context, dbtx *sqlx.Tx, req *models.Bank) error
	UpdateBank(ctx context.Context, dbtx *sqlx.Tx, req *models.Bank) error
	DeleteBank(ctx context.Context, dbtx *sqlx.Tx, bankID int) error
}