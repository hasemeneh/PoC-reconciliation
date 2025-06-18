package transaction

import (
	"context"

	"github.com/hasemeneh/PoC-OnlineStore/helper/database"
	"github.com/hasemeneh/PoC-reconciliation/svc/reconcileapp/src/models"
	"github.com/jmoiron/sqlx"
)

var getTransactionByIDQuery = "SELECT `trxID`, `amount`, `type`, `transactionTime`, `bankID`, `isReconciled` FROM `Transaction` WHERE `trxID` = ?;"
var insertTransactionQuery = "INSERT INTO `Transaction` (`trxID`, `amount`, `type`, `transactionTime`, `isReconciled`, `bankID`) VALUES (?, ?, ?, ?, ?, ?);"
var getAllTransactionsQuery = "SELECT `trxID`, `amount`, `type`, `transactionTime`, `bankID`, `isReconciled` FROM `Transaction`;"
var updateTransactionQuery = "UPDATE `Transaction` SET `amount` = ?, `type` = ?, `transactionTime` = ?, `isReconciled` = ?, `bankID` = ? WHERE `trxID` = ?;"
var deleteTransactionQuery = "DELETE FROM `Transaction` WHERE `trxID` = ?;"

func (p *transactionDomain) GetTransactionByID(ctx context.Context, trxID string) (*models.TransactionModel, error) {
	var transaction models.TransactionModel
	err := p.DB.GetContext(ctx, &transaction, getTransactionByIDQuery, trxID)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (p *transactionDomain) GetAllTransactions(ctx context.Context) ([]*models.TransactionModel, error) {
	var transactions []*models.TransactionModel
	err := p.DB.SelectContext(ctx, &transactions, getAllTransactionsQuery)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (p *transactionDomain) UpdateTransaction(ctx context.Context, dbtx *sqlx.Tx, req *models.TransactionModel) error {
	var execFunc database.ExecFunc
	if dbtx == nil {
		execFunc = p.DB.ExecContext
	} else {
		execFunc = dbtx.ExecContext
	}
	_, err := execFunc(ctx, updateTransactionQuery, req.Amount, req.Type, req.TransactionTime, req.IsReconciled, req.BankID, req.TrxID)
	return err
}

func (p *transactionDomain) DeleteTransaction(ctx context.Context, dbtx *sqlx.Tx, trxID string) error {
	var execFunc database.ExecFunc
	if dbtx == nil {
		execFunc = p.DB.ExecContext
	} else {
		execFunc = dbtx.ExecContext
	}
	_, err := execFunc(ctx, deleteTransactionQuery, trxID)
	return err
}

func (p *transactionDomain) InsertTransaction(ctx context.Context, dbtx *sqlx.Tx, req *models.TransactionModel) error {
	var execFunc database.ExecFunc
	if dbtx == nil {
		execFunc = p.DB.ExecContext
	} else {
		execFunc = dbtx.ExecContext
	}
	_, err := execFunc(ctx, insertTransactionQuery, req.TrxID, req.Amount, req.Type, req.TransactionTime, req.IsReconciled, req.BankID)
	return err
}

func (p *transactionDomain) StartTx(ctx context.Context) (*sqlx.Tx, error) {
	return p.DB.BeginTxx(ctx, nil)
}
