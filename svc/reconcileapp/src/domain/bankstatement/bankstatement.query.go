package bankstatement

import (
	"context"

	"github.com/hasemeneh/PoC-OnlineStore/helper/database"
	"github.com/hasemeneh/PoC-reconciliation/svc/reconcileapp/src/models"
	"github.com/jmoiron/sqlx"
)

var getStatementByIDQuery = "SELECT `unique_identifier`, `amount`, `statement_date`, `bankID`, `isReconciled` FROM `BankStatement` WHERE `unique_identifier` = ?;"
var insertStatementQuery = "INSERT INTO `BankStatement` (`unique_identifier`, `amount`, `statement_date`, `bankID`, `isReconciled`) VALUES (?, ?, ?, ?, ?);"
var getAllStatementsQuery = "SELECT `unique_identifier`, `amount`, `statement_date`, `bankID`, `isReconciled` FROM `BankStatement`;"
var updateStatementQuery = "UPDATE `BankStatement` SET `amount` = ?, `statement_date` = ?, `isReconciled` = ? WHERE `unique_identifier` = ? AND `bankID` = ?;"
var deleteStatementQuery = "DELETE FROM `BankStatement` WHERE `unique_identifier` = ?;"

func (p *bankstatementDomain) GetStatementByID(ctx context.Context, uniqueIdentifier string) (*models.BankStatement, error) {
	var statement models.BankStatement
	err := p.DB.GetContext(ctx, &statement, getStatementByIDQuery, uniqueIdentifier)
	if err != nil {
		return nil, err
	}
	return &statement, nil
}

func (p *bankstatementDomain) GetAllStatements(ctx context.Context) ([]*models.BankStatement, error) {
	var statements []*models.BankStatement
	err := p.DB.SelectContext(ctx, &statements, getAllStatementsQuery)
	if err != nil {
		return nil, err
	}
	return statements, nil
}

func (p *bankstatementDomain) UpdateStatement(ctx context.Context, dbtx *sqlx.Tx, req *models.BankStatement) error {
	var execFunc database.ExecFunc
	if dbtx == nil {
		execFunc = p.DB.ExecContext
	} else {
		execFunc = dbtx.ExecContext
	}
	_, err := execFunc(ctx, updateStatementQuery, req.Amount, req.Date, req.IsReconciled, req.UniqueIdentifier, req.BankID)
	return err
}

func (p *bankstatementDomain) DeleteStatement(ctx context.Context, dbtx *sqlx.Tx, uniqueIdentifier string) error {
	var execFunc database.ExecFunc
	if dbtx == nil {
		execFunc = p.DB.ExecContext
	} else {
		execFunc = dbtx.ExecContext
	}
	_, err := execFunc(ctx, deleteStatementQuery, uniqueIdentifier)
	return err
}

func (p *bankstatementDomain) InsertStatement(ctx context.Context, dbtx *sqlx.Tx, req *models.BankStatement) error {
	var execFunc database.ExecFunc
	if dbtx == nil {
		execFunc = p.DB.ExecContext
	} else {
		execFunc = dbtx.ExecContext
	}
	_, err := execFunc(ctx, insertStatementQuery, req.UniqueIdentifier, req.Amount, req.Date, req.BankID, req.IsReconciled)
	return err
}

func (p *bankstatementDomain) StartTx(ctx context.Context) (*sqlx.Tx, error) {
	return p.DB.BeginTxx(ctx, nil)
}
